package easy_config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Viper *viper.Viper

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

	v.WatchConfig()
	// 判断config结构体
	if err = v.Unmarshal(&Config); err != nil {
		fmt.Println(err)
	}
	// 监听改变
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&Config); err != nil {
			fmt.Println(err)
		}
	})
	Viper = v
	return v
}
