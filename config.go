package github.com/yicheng20110203/locker

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"sync"
)

type Config struct {
	Etcd      configEtcd      `yaml:"etcd"`
	Redis     configRedis     `yaml:"redis"`
	Zookeeper configZookeeper `yaml:"zookeeper"`
	Consule   configConfule   `yaml:"consule"`
}

type configEtcd struct {
	Servers   []string `yaml:"servers"`
	DailTimes int64    `yaml:"dailtimes"`
}

type configRedis struct {
	Servers   []string `yaml:"servers"`
	DailTimes int64    `yaml:"dailtimes"`
}

type configZookeeper struct {
	Servers   []string `yaml:"servers"`
	DailTimes int64    `yaml:"dailtimes"`
}

type configConfule struct {
	Servers   []string `yaml:"servers"`
	DailTimes int64    `yaml:"dailtimes"`
}

var (
	once sync.Once
	Cfg  *Config
)

func LoadConfig() (err error) {
	once.Do(func() {
		var path string
		path, err = os.Getwd()
		if err != nil {
			log.Errorf("LoadConfig().os.Getwd() error: %+v", err)
			return
		}

		cfg := viper.New()
		cfg.SetConfigName("config")
		cfg.SetConfigType("yml")
		cfg.AddConfigPath(path)
		err = cfg.ReadInConfig()
		if err != nil {
			log.Errorf("LoadConfig() cfg.ReadInConfig() error: %+v", err)
			return
		}

		var readerCfg = &Config{}
		if err = cfg.Unmarshal(&readerCfg); err != nil {
			log.Errorf("LoadConfig() cfg.Unmarshal(Cfg) error: %+v", err)
			return
		}

		Cfg = readerCfg
		return
	})

	return
}
