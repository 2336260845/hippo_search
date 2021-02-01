package main

import (
	"encoding/json"
	"fmt"
	"github.com/2336260845/hippo_search/client"
	"github.com/2336260845/hippo_search/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ginEngineInit() *gin.Engine {
	r := gin.Default()
	register(r)
	return r
}

func configInit() *config.Config {
	config.InitConfig()

	cf := config.GetConfig()

	if cf.HttpAddress == "" {
		cf.HttpAddress = "0.0.0.0:8999"
	}

	b, _ := json.Marshal(cf)
	logrus.Infof("configInit config=%+v", string(b))

	return cf
}

func main() {
	r := ginEngineInit()
	conf := configInit()

	client.ThriftInit(conf)

	if err := r.Run(conf.HttpAddress); err != nil {
		panic(fmt.Sprintf("server run error, err=%s", err.Error()))
	}
}
