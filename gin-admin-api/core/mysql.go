// mysql初始化配置
// @author chen

package core

import (
	"fmt"
	"gin-admin-api/config"
	"gin-admin-api/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func MysqlInit() error {
	var err error
	var dbConfig = config.Config.Mysql

	// 尝试连接数据库
	tempUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Charset)

	tempDB, err := gorm.Open(mysql.Open(tempUrl), &gorm.Config{})
	if err != nil {
		return err
	}

	// 创建数据库
	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", dbConfig.Db)
	if err := tempDB.Exec(createDBSQL).Error; err != nil {
		return fmt.Errorf("创建数据库失败:%v", err)
	}

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
		dbConfig.Charset)
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
