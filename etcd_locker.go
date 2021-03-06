package locker

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	log "github.com/sirupsen/logrus"
)

type EtcdLocker struct {
	Ttl     int64
	Conf    clientv3.Config
	Key     string
	cancel  context.CancelFunc
	lease   clientv3.Lease
	leaseId clientv3.LeaseID
	txn     clientv3.Txn
}

func (em *EtcdLocker) init() (err error) {
	var client *clientv3.Client
	client, err = clientv3.New(em.Conf)
	if err != nil {
		log.Errorf("EtcdLocker.init() clientv3.New error: %+v", err)
		return
	}

	em.lease = clientv3.NewLease(client)
	var leaseResp *clientv3.LeaseGrantResponse
	leaseResp, err = em.lease.Grant(context.TODO(), em.Ttl)
	if err != nil {
		log.Errorf("EtcdLocker.init() em.lease.Grant error: %+v", err)
		return
	}
	em.leaseId = leaseResp.ID
	em.txn = clientv3.NewKV(client).Txn(context.TODO())
	var ctx context.Context
	ctx, em.cancel = context.WithCancel(context.TODO())
	_, err = em.lease.KeepAlive(ctx, em.leaseId)
	if err != nil {
		log.Errorf("EtcdLocker.init() em.lease.KeepAlive error: %+v", err)
		return
	}
	return
}

func (em *EtcdLocker) Lock() (err error) {
	err = em.init()
	if err != nil {
		return
	}

	em.txn.If(clientv3.Compare(clientv3.CreateRevision(em.Key), "=", 0)).Then(clientv3.OpPut(em.Key, "", clientv3.WithLease(em.leaseId))).Else(clientv3.OpGet(em.Key))
	var txnResp *clientv3.TxnResponse
	txnResp, err = em.txn.Commit()
	if err != nil {
		log.Errorf("EtcdLocker.Lock() em.txn.Commit() error: %+v", err)
		return
	}

	if !txnResp.Succeeded {
		log.Info("EtcdLocker.Lock() failed")
		err = ErrorLockerHasBeenUsed
		return
	}
	return
}

func (em *EtcdLocker) Unlock() {
	var err error
	em.cancel()
	_, err = em.lease.Revoke(context.TODO(), em.leaseId)
	if err != nil {
		log.Errorf("EtcdLocker.Unlock() em.lease.Revoke error: %+v", err)
		return
	}
}