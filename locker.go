package locker

import (
	"github.com/coreos/etcd/clientv3"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type ILocker interface {
	Lock() error
	Unlock()
}

type Locker struct {
	lockType    int
	serverList  []string
	dailTimeout int64
	key         string
	ttl         int64
	mutex       ILocker
}

func NewLocker(lockType int) (lock *Locker) {
	lock = &Locker{
		lockType:    lockType,
		serverList:  ServerListFromCfg(lockType),
		dailTimeout: DailTimeoutFromCfg(lockType),
	}
	return
}

func (lock *Locker) WithKey(key string, ttl int64) {
	lock.key = key
	lock.ttl = ttl
}

var (
	mut sync.Mutex
)

func (lock *Locker) Lock() (err error) {
	mut.Lock()
	defer mut.Unlock()

	switch lock.lockType {
	case LockTypeEtcd:
		mt := &EtcdLocker{
			Ttl: lock.ttl,
			Conf: clientv3.Config{
				Endpoints:   lock.serverList,
				DialTimeout: time.Duration(lock.dailTimeout) * time.Second,
			},
			Key: lock.key,
		}
		err = mt.init()
		if err != nil {
			log.Errorf("Locker.Lock() mt.init() error: %+v", err)
			return
		}
		lock.mutex = mt
		err = mt.Lock()
		return
	case LockTypeRedis:
		mt := &RedisLocker{
			Ttl:    lock.ttl,
			Key:    lock.key,
			Server: lock.serverList,
		}
		mt.init()
		lock.mutex = mt
		err = mt.Lock()
		return
	}

	return
}

func (lock *Locker) Unlock() {
	switch lock.lockType {
	case LockTypeEtcd:
		lock.mutex.Unlock()
		return
	}
}
