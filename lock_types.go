package locker

const (
	LockTypeZk = 1 << iota
	LockTypeEtcd
	LockTypeRedis
	LockTypeConsule
)