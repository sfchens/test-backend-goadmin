package apis

import (
	"csf/app/admin/service/sys_service"
	"csf/library/easy_config"
	"csf/library/easy_db"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	configYml  string
	AppRouters = make([]func(), 0)
	StartCmd   = &cobra.Command{
		Use:          "initApi",
		Short:        "Start initApi",
		Example:      "main initApi -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			// 初始化
			setUp()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

// 配置类
func setUp() {
	var err error
	// 初始化Config
	easy_config.InitConfig()
	fmt.Println("viper config 初始化成功")

	// 初始化数据库
	err = easy_db.InitMysql("mysql")
	if err != nil {
		fmt.Printf("database 数据库初始化失败； 错误信息： %v\n", err.Error())
	}
	fmt.Println("mysql 初始化成功")
}

func run() (err error) {
	// 初始化admin路由入库
	initRegisterRouter()

	return
}

func initRegisterRouter() {
	if easy_config.Config.App.IsApiMysql {
		ctx := &gin.Context{}
		err := sys_service.NewSysApiService(ctx).Refresh()
		if err != nil {
			println("初始化Api数据失败")
		}
	}
}
