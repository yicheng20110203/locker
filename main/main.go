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

	log.Infof("read yaml config: %+v", locker.Cfg)

	TT()
	time.Sleep(time.Duration(5) * time.Second)
}


func TT() {
	go func() {
		lock1 := locker.NewLocker(locker.LockTypeEtcd)
		lock1.WithKey("dy", 1)
		err := lock1.Lock()
		defer lock1.Unlock()
		time.Sleep(2 * time.Second)
		if err != nil {
			log.Errorf("go 1 lock error: %+v", err)
			return
		}
		log.Info("go 1 lock success")
	}()

	go func() {
		lock1 := locker.NewLocker(locker.LockTypeEtcd)
		lock1.WithKey("dy", 1)
		err := lock1.Lock()
		defer lock1.Unlock()
		time.Sleep(2 * time.Second)
		if err != nil {
			log.Errorf("go 2 lock error: %+v", err)
			return
		}
		log.Info("go 2 lock success")
	}()
}