package github.com/yicheng20110203/locker

const (
	LockTypeZk = 1 << iota
	LockTypeEtcd
	LockTypeRedis
	LockTypeConsule
)