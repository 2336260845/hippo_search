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
	HasParse     bool          `json:"-"`
	Version      string        `json:"version"`
	HttpAddress  string        `json:"httpAddress"`
	ServerConfig *ServerConfig `json:"server_config"` //TODO 有问题，下划线命名不识别
	TimeCircle   TimeCircle
	SkipModule   SkipModule
}

type ServerConfig struct {
	QueryAddress  string `json:"queryAddress"`
	RecallAddress string `json:"recallAddress"`
	RankAddress   string `json:"rankAddress"`
}

type TimeCircle struct {
	QueryCut int
}

type SkipModule struct {
	SkipQueryAnalysis bool
	SkipRecall        bool
	SkipRank          bool
}

var conf *Config

func InitConfig() {
	if conf != nil && conf.HasParse {
		logrus.Warnf("InitConfig has parse config, config=%+v", conf)
		return
	}

	cf := parseConfig()
	if err := checkConfig(cf); err != nil {
		panic(fmt.Sprintf("InitConfig checkConfig error, err=%+v", err.Error()))
	}

	conf = cf
	cf.HasParse = true
}

func GetConfig() *Config {
	return conf
}

func parseConfig() *Config {
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

func checkConfig(conf *Config) error {
	if conf.TimeCircle.QueryCut < 5 {
		return fmt.Errorf("query cut time is too little, time=%+v", conf.TimeCircle.QueryCut)
	}

	return nil
}
