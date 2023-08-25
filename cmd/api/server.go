package api

import (
	"csf/app/admin/router"
	"csf/library/easy_config"
	"csf/library/easy_db"
	"csf/library/easy_logger"
	"csf/library/easy_validator"
	"csf/library/global"
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
	// 初始化Config
	easy_config.InitConfig()
	fmt.Println("viper config 初始化成功")

	// 初始化数据库
	err = easy_db.InitMysql("mysql")
	if err != nil {
		fmt.Printf("database 数据库初始化失败； 错误信息： %v\n", err.Error())
	}
	fmt.Println("mysql 初始化成功")

	// 加载验证器
	easy_validator.InitValidate()
	fmt.Println("验证器 初始化成功")

	// 初始化日志
	easy_logger.InitLogger()
	fmt.Println("日志 初始化成功")
}

func run() (err error) {

	if easy_config.Viper.GetString("app.mode") == global.ModeProd {
		//gin.SetMode(gin.ReleaseMode)
	}
	// 加载公共中间件路由
	initApiRouter()
	// 加载路由
	for _, f := range AppRouters {
		f()
	}

	// 初始化路由入库
	//initRegisterRouter()

	// 运行
	err = endless.ListenAndServe(fmt.Sprintf(":%d", easy_config.Viper.Get("app.port")), global.GinEngine)
	if err != nil {
		fmt.Printf("运行失败:  %+v\n", err.Error())
	}
	return nil
}
