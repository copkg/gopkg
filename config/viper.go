package config

import (
	"github.com/spf13/viper"
	"log"
)

var Conf *viper.Viper

func MustLoad(path string) *viper.Viper {
	var err error
	conf := viper.New()
	conf.SetConfigFile(path)
	if err = conf.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Panicf("config file not found")
		} else {
			log.Panicf("read config file failed:%s", err.Error())
		}
	}
	log.Println("config file load success")
	return conf
}
