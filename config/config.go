package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	ConfigPath = "./script/"
)

type Config struct {
	Version      string       `json:"version"`
	Address      string       `json:"address"`
	ServerConfig ServerConfig `json:"server_config"` //TODO 有问题，结构体无法解析
}

type ServerConfig struct {
	Address string `json:"address"`
}

func ParseConfig() *Config {
	cf := Config{}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(ConfigPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("ParseConfig path:%s read error, err=%s", ConfigPath, err.Error()))
	}

	err = viper.Unmarshal(&cf)
	if err != nil {
		panic(fmt.Sprintf("ParseConfig path:%s unmarshl error, err=%s", ConfigPath, err.Error()))
	}

	logrus.Infof("ParseConfig config=%+v", cf)
	return &cf
}
