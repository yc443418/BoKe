// mysql初始化配置
// @author chen
package core

import (
	"admin_vue3/config"
	"admin_vue3/global"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func MysqlInit() error {
	var err error
	var dbConfig = config.Config.Mysql
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Db, dbConfig.Charset)

	Db, err = gorm.Open(mysql.Open(url), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	if Db.Error != nil {
		return err
	}

	sqlDb, err := Db.DB()
	sqlDb.SetMaxIdleConns(dbConfig.MaxIdle)
	sqlDb.SetMaxOpenConns(dbConfig.MaxOpen)
	global.Log.Infof("[mysql]连接成功")
	return nil
}
