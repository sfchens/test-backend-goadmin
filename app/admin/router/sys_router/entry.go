package sys_router

import (
	"csf/library/global"
	"fmt"
)

func init() {
	global.RouterList = append(global.RouterList,
		registerLoginRouter,
		registerConfigRouter,
		registerAdminRouter,
		registerMenuRouter,
		registerDeptRouter,
		registerSwitchRouter,
		registerApiRouter,
		registerRoleRouter,
	)

}

func InitRouter() {
	fmt.Printf("初始化 sys_router路由")
}
