package db

import (
	"csf/library/viper"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// gormMysql 初始化数据库
func gormMysql(mysqlName string) error {
	mysqlKey := fmt.Sprintf("database.%s", mysqlName)

	mysqlConfig := mysql.Config{
		DSN:                       Dsn(mysqlName),
		DefaultStringSize:         256,   //string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   viper.NewViper.GetString(mysqlKey + ".tablePrefix"), // 表前缀
			SingularTable: true,
		},
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), config)
	if err != nil {
		return err
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(viper.NewViper.GetInt(mysqlKey + ".maxIdConnect"))
		sqlDB.SetMaxOpenConns(viper.NewViper.GetInt(mysqlKey + ".maxOpenConnect"))
		newMysql[mysqlName] = db
	}
	return nil
}

func Dsn(mysqlName string) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&%s",
		viper.NewViper.Get(fmt.Sprintf("database.%s.username", mysqlName)),
		viper.NewViper.Get(fmt.Sprintf("database.%s.password", mysqlName)),
		viper.NewViper.Get(fmt.Sprintf("database.%s.host", mysqlName)),
		viper.NewViper.Get(fmt.Sprintf("database.%s.port", mysqlName)),
		viper.NewViper.Get(fmt.Sprintf("database.%s.dbname", mysqlName)),
		viper.NewViper.Get(fmt.Sprintf("database.%s.charset", mysqlName)),
		viper.NewViper.Get(fmt.Sprintf("database.%s.extra", mysqlName)),
	)
	return dsn
}
