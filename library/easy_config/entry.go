package easy_config

import (
	"csf/configs/config"
	"fmt"
)

var Config config.Server

func InitConfig(path ...string) {
	fmt.Printf("config: %+v\n", Config)
	// 初始化Viper
	InitViper(path...)
}
