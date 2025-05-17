package middleware

import (
	"gin-admin-api/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.BArticle{},        // 文章模型
		&model.BArticleComment{}, // 文章评论
		&model.BCarousel{},       // 轮播图
		&model.BCategory{},       // 文章分类
		&model.BTags{},           // 标签
		&model.SysAdmin{},        // 用户
		&model.SysAdminRole{},    // 用户与角色关系
		&model.SysConfig{},       // 系统参数
		&model.SysMenu{},         // 菜单列表
		&model.SysRole{},         // 角色
		&model.SysRoleMenu{},     // 角色与菜单
		&model.User{},            // 前台用户
		&model.UserMessage{},     // 消息
	)
}
