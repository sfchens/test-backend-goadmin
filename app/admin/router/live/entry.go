package live

import (
	"csf/library/global"
	"fmt"
)

func InitRouter() {
	fmt.Println("初始化Common路由")
	global.RouterList[global.ModuleAdmin] = append(global.RouterList[global.ModuleAdmin],
		registerBackdropRouter,
		registerVideoRouter,
	)
}
