package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/yicheng20110203/locker"
	"time"
)

func main() {
	var err error
	if err = locker.LoadConfig(); err != nil {
		log.Fatal(err)
		return
	}

	// ETCD locker
	etcd()

	// redis locker
	redis()

	time.Sleep(time.Duration(5) * time.Second)
}


func etcd() {
	go func() {
		lock1 := locker.NewLocker(locker.LockTypeEtcd)
		lock1.WithKey("dy", 1)
		err := lock1.Lock()
		defer lock1.Unlock()
		time.Sleep(2 * time.Second)
		if err != nil {
			log.Errorf("etcd: go 1 lock error: %+v", err)
			return
		}
		log.Info("etcd: go 1 lock success")
	}()

	go func() {
		lock1 := locker.NewLocker(locker.LockTypeEtcd)
		lock1.WithKey("dy", 1)
		err := lock1.Lock()
		defer lock1.Unlock()
		time.Sleep(2 * time.Second)
		if err != nil {
			log.Errorf("etcd: go 2 lock error: %+v", err)
			return
		}
		log.Info("etcd: go 2 lock success")
	}()
}

func redis() {
	go func() {
		lock1 := locker.NewLocker(locker.LockTypeRedis)
		lock1.WithKey("dy", 10)
		err := lock1.Lock()
		defer lock1.Unlock()
		time.Sleep(2 * time.Second)
		if err != nil {
			log.Errorf("redis: go 1 lock error: %+v", err)
			return
		}
		log.Info("redis: go 1 lock success")
	}()

	go func() {
		lock1 := locker.NewLocker(locker.LockTypeRedis)
		lock1.WithKey("dy", 10)
		err := lock1.Lock()
		defer lock1.Unlock()
		time.Sleep(2 * time.Second)
		if err != nil {
			log.Errorf("redis: go 2 lock error: %+v", err)
			return
		}
		log.Info("redis: go 2 lock success")
	}()
}