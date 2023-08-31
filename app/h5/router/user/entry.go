package user

import (
	"csf/library/global"
	"fmt"
)

func InitRouter() {
	fmt.Println("初始化User路由")
	global.RouterList[global.ModuleH5] = append(global.RouterList[global.ModuleH5],
		registerUserRouter,
	)
}
