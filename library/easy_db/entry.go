package easy_db

import (
	"csf/library/easy_config"
	"fmt"
	"gorm.io/gorm"
)

var newMysql = make(map[string]*gorm.DB, 1)

const defaultMysqlName = "mysql"

func GetDb(dbName ...string) (db *gorm.DB) {
	if len(dbName) > 1 {
		return
	}
	if len(dbName) == 1 {
		db = newMysql[dbName[0]]
		return
	}
	db = newMysql[defaultMysqlName]
	return
}

func InitMysql(mysqlName string) error {
	// 根据驱动配置进行初始化
	driverType := easy_config.Viper.GetString(fmt.Sprintf("database.%s.driver", mysqlName))
	switch driverType {
	case "mysql":
		return gormMysql(mysqlName)
	default:
		return gormMysql(mysqlName)
	}
}
