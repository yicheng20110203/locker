package locker

func ServerListFromCfg(lockerType int) []string {
	switch lockerType {
	case LockTypeEtcd:
		return Cfg.Etcd.Servers
	case LockTypeRedis:
		return Cfg.Redis.Servers
	case LockTypeZk:
		return Cfg.Zookeeper.Servers
	}

	return nil
}

func DailTimeoutFromCfg(lockerType int) int64 {
	switch lockerType {
	case LockTypeEtcd:
		return Cfg.Etcd.DailTimes
	case LockTypeRedis:
		return Cfg.Redis.DailTimes
	case LockTypeZk:
		return Cfg.Zookeeper.DailTimes
	}

	return 0
}
