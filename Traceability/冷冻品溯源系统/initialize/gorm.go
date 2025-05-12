package initialize

import (
	"Traceability/global"
	"Traceability/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Factory{},
		&model.Dealer{},
		&model.Consumer{},
		&model.Supervision{},
		&model.FrozenProduct{},
		&model.Logistics{},
		&model.Order{},
		&model.AuditLog{},
		&model.ProductApplication{},
		//中间表
		//&model.ConsumerProduct{},
		//&model.SupervisionProduct{},
	)
}

func MustLoadGorm() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", global.CONFIG.MySQL.Username, global.CONFIG.MySQL.Password,
		global.CONFIG.MySQL.Host, global.CONFIG.MySQL.Port, global.CONFIG.MySQL.Database, global.CONFIG.MySQL.Charset)
	fmt.Println("我的mysql", dsn)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	global.DB = db
	MysqlDB, _ := db.DB()
	MysqlDB.SetMaxIdleConns(10)
	MysqlDB.SetMaxOpenConns(20)

}
