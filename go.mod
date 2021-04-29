module github.com/yicheng20110203/locker

go 1.14

require (
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/go-redis/redis/v8 v8.8.2
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.7.1
	go.uber.org/zap v1.16.0 // indirect
	google.golang.org/genproto v0.0.0-20210405174219-a39eb2f71cb9 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
