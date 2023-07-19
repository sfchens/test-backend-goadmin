package viper

import (
	"fmt"
	"github.com/spf13/viper"
)

var NewViper *viper.Viper

func InitViper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		config = "./configs/config.yaml"
	} else { // 函数传递的可变参数的第一个值赋值于config
		config = path[0]
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	NewViper = v
	return v
}
