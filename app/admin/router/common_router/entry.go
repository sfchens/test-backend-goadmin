package common_router

import (
	"csf/library/global"
	"fmt"
)

func init() {
	global.RouterList = append(global.RouterList,
		registerOtherRouter,
		registerCommonRouter,
		registerTestRouter,
		registerUploadRouter,
	)
}
func InitRouter() {
	fmt.Printf("初始化Common路由")
}
