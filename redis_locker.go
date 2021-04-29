package locker

import (
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"time"
)

type RedisLocker struct {
	Ttl    int64
	Server []string
	Key    string
	rdb    *redis.Client
	ctx    context.Context
}

func (rl *RedisLocker) init() {
	rl.rdb = redis.NewClient(&redis.Options{
		Addr:     rl.Server[0],
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rl.ctx = context.Background()
}

func (rl *RedisLocker) Lock() (err error) {
	var setRes bool
	setRes, err = rl.rdb.SetNX(rl.ctx, rl.Key, 1, time.Duration(rl.Ttl)*time.Second).Result()
	if err != nil {
		log.Errorf("RedisLocker.Lock() rl.rdb.SetNX error: %+v", err)
		return
	}

	if !setRes {
		err = ErrorLockerHasBeenUsed
		return
	}

	return
}

func (rl *RedisLocker) Unlock() {
	nums, err := rl.rdb.Del(rl.ctx, rl.Key).Result()
	if err != nil {
		log.Errorf("RedisLocker.Unlock() error: %+v", err)
		return
	}

	if nums <= 0 {
		log.Errorf("RedisLocker.Unlock() failed")
	}
}
