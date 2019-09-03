package main

import (
	"github.com/sirupsen/logrus"
	"hermes/cmd"
	"hermes/config"
	"hermes/pkg/drivers/redis"
)

func main() {
	err := config.FromEnv(nil)
	if err != nil {
		logrus.Fatalf("error in loading config map from env: %v", err)
	}
	cmd.Launch()
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("panic happend:\n%v", err)
		}
		con, err := redis.ConnectRedis()
		if err != nil {
			logrus.Errorf("error while trying to clear redis cache of subscribed channels : %v", err)
			return
		}
		_, err = con.FlushDB().Result()
		if err != nil {
			logrus.Errorf("error while trying to clear redis cache of subscribed channels : %v", err)
			return
		}
	}()
}
