package easy_config

import (
	"csf/configs/config"
)

var Config config.Server

func InitConfig(path ...string) {
	// 初始化Viper
	InitViper(path...)
}
