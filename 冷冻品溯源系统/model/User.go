package model

import (
	"Traceability/utils"
	"gorm.io/gorm"
	"time"
)

// 用户表
type User struct {
	UserID      string         `gorm:"primaryKey;type:varchar(36)" json:"user_id"`
	Username    string         `gorm:"type:varchar(50);unique;not null" json:"username" binding:"required,min=5,max=50"`
	Password    string         `gorm:"type:varchar(100);not null" json:"password" binding:"required,min=8,max=50"`
	UserType    string         `gorm:"type:ENUM('factory','dealer','supervision','consumer','admin');index;not null" json:"user_type" binding:"required,oneof=factory dealer supervision consumer admin"`
	Status      int            `gorm:"type:int;default:1;index;comment:1启用 2禁用" json:"status"`
	ContactInfo Contact        `gorm:"embedded;not null" json:"contact_info" binding:"required"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// 用户类型：factory(工厂)/dealer(经销商)/supervision(监管)/consumer(消费者)/admin(管理员)

// Contact 联系信息结构体
type Contact struct {
	Name  string `gorm:"type:varchar(50);comment:联系人姓名" json:"name" binding:"required"`
	Email string `gorm:"type:varchar(100);comment:电子邮箱地址" json:"email" binding:"required,email"`
	Phone string `gorm:"type:varchar(20);comment:联系电话号码" json:"phone" binding:"required"`
}

// BeforeCreate 创建用户前生成UUID并加密密码
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.UserID = utils.GenerateUUID()

	// 密码加密
	if hashedPwd, err := utils.HashPassword(u.Password); err != nil {
		return err
	} else {
		u.Password = hashedPwd
	}
	return nil
}
