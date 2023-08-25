package logic

import (
	_ "csf/core/logic/common"
	_ "csf/core/logic/config"
	_ "csf/core/logic/live"
	_ "csf/core/logic/login"
	_ "csf/core/logic/sys"
	_ "csf/core/logic/user"
	"fmt"
)

func InitService() {
	fmt.Printf("初始化service")
}
