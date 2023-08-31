package common

import (
	"csf/library/global"
)

func InitRouter() {
	global.RouterList[global.ModuleAdmin] = append(global.RouterList[global.ModuleAdmin],
		registerOtherRouter,
		registerCommonRouter,
		registerTestRouter,
		registerUploadRouter,
	)
}
