package api

import (
	"csf/app/admin/router"
	"csf/library/db"
	"csf/library/global"
	"csf/library/my_validator"
	"csf/library/viper"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/spf13/cobra"
)

var (
	configYml  string
	AppRouters = make([]func(), 0)
	StartCmd   = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "main server -c config/settings.yml",
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
	AppRouters = append(AppRouters, router.InitRouter)
}

// 配置类
func setUp() {
	var err error
	// 初始化viper
	viper.InitViper()
	// 初始化数据库
	err = db.InitMysql("mysql")
	if err != nil {
		fmt.Printf("database 数据库初始化失败； 错误信息： %v\n", err.Error())
	}
	// 加载验证器
	my_validator.InitValidate()
	fmt.Printf("mysql 初始化成功")
}

func run() (err error) {

	if viper.NewViper.GetString("app.mode") == global.ModeProd {
		//gin.SetMode(gin.ReleaseMode)
	}
	// 加载Api路由
	initApiRouter()
	for _, f := range AppRouters {
		f()
	}

	// 初始化路由入库
	initRegisterRouter()
	// 运行
	err = endless.ListenAndServe(fmt.Sprintf(":%d", viper.NewViper.Get("app.port")), global.GinEngine)
	if err != nil {
		fmt.Printf("运行失败:  %+v\n", err.Error())
	}
	return nil
}
