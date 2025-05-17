package middleware

import (
	"gin-front-api/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.BArticle{},            // 文章模型
		&model.BArticleComment{},     // 文章评论
		&model.BCarousel{},           // 轮播图
		&model.BCategory{},           // 文章分类
		&model.BTags{},               // 标签
		&model.BArticleCommentStar{}, // 文章评论点赞
		&model.BArticleStar{},        // 文章点赞
		&model.SysConfig{},           // 系统参数
		&model.User{},                // 前台用户
		&model.UserMessage{},         // 消息
	)
}
