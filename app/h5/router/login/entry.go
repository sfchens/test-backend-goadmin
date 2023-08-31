package login

import "csf/library/global"

func InitRouter() {
	global.RouterList[global.ModuleH5] = append(global.RouterList[global.ModuleH5],
		registerLoginRouter,
	)
}
