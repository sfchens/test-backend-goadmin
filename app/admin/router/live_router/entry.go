package live_router

import (
	"csf/library/global"
	"fmt"
)

func init() {
	global.RouterList = append(global.RouterList,
		registerBackdropRouter,
		registerVideoRouter,
	)
}

func InitRouter() {
	fmt.Printf("初始化Common路由")
}
