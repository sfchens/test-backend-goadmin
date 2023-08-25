package user

import (
	"csf/library/global"
	"fmt"
)

func InitRouter() {
	fmt.Println("初始化User路由")
	global.RouterList = append(global.RouterList,
		registerUserRouter,
	)
}
