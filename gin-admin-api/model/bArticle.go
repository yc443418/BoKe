// 文章相关模型
// @author chen

package model

import util "gin-admin-api/utils"

// BArticle 文章模型对象
type BArticle struct {
	ID              uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                        // ID
	UserId          uint       `gorm:"column:user_id;comment:'用户id';NOT NULL" json:"userId"`                        // 用户id
	CategoryId      uint       `gorm:"column:category_id;comment:'分类id';NOT NULL" json:"categoryId"`                // 分类id
	TagId           uint       `gorm:"column:tag_id;comment:'标签id';NOT NULL" json:"tagId"`                          // 标签id
	Title           string     `gorm:"column:title;varchar(150);comment:'标题'" json:"title"`                         //  文章标题
	Image           string     `gorm:"column:image;varchar(200);comment:'封面地址'" json:"image"`                       //  封面地址
	Summary         string     `gorm:"column:summary;varchar(255);comment:'文章简介'" json:"summary"`                   //  文章简介
	MarkdownContent string     `gorm:"column:markdown_content;comment:'文章内容md版'" json:"markdownContent"`            //  文章内容md版
	TopType         uint       `gorm:"column:top_type;comment:'是否发布(1-未发布 2-已发布)';DEFAULT 1" json:"topType"`        // 是否置顶(1-未置顶 2-已置顶)
	Quantity        uint       `gorm:"column:quantity;default:0;comment:'标签id';NOT NULL" json:"quantity"`           // 文章阅读量
	GoodCount       uint       `gorm:"column:good_count;default:0;comment:'文章点赞量';NOT NULL" json:"goodCount"`       // 文章点赞量
	CommentCount    uint       `gorm:"column:comment_count;default:0;comment:'文章评论量';NOT NULL" json:"commentCount"` // 文章评论量
	CreateTime      util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`                // 创建时间
	Status          uint       `gorm:"column:status; comment:'状态（1-已删除，2-未删除）';DEFAULT 2" json:"status"`            // 状态（1-已删除，2-未删除）
}

func (BArticle) TableName() string {
	return "b_article"
}

// BArticleVo 文章视图对象
type BArticleVo struct {
	ID           uint       `json:"id"`           // ID
	Image        string     `json:"image"`        // 封面地址
	Title        string     `json:"title"`        // 文章标题
	CategoryId   uint       `json:"categoryId"`   // 分类id
	CategoryName string     `json:"categoryName"` // 分类名称
	TagId        uint       `json:"tagId"`        // 标签id
	TagName      string     `json:"tagName"`      // 标签名称
	UserId       uint       `json:"userId"`       // 用户id
	Username     string     `json:"username"`     // 用户名称
	Summary      string     `json:"summary"`      // 文章简介
	TopType      uint       `json:"topType"`      // 是否置顶(1-未置顶 2-已置顶)
	Quantity     uint       `json:"quantity"`     // 文章阅读量
	GoodCount    uint       `json:"goodCount"`    // 文章点赞量
	CommentCount uint       `json:"commentCount"` // 文章评论量
	CreateTime   util.HTime `json:"createTime"`   // 创建时间
	Status       uint       `json:"status"`       // 状态（1-已删除，2-未删除）
}

// UpdateTopTypeDto 设置置顶
type UpdateTopTypeDto struct {
	ID      uint `json:"id"`      // ID
	TopType uint `json:"topType"` // 是否置顶(1-未置顶 2-已置顶)
}

// BArticleIdDto 文章id
type BArticleIdDto struct {
	ID int `json:"id"` // ID
}

// CommentVo 文章评论对象
type CommentVo struct {
	Content      string     `json:"content"`      // 评论内容
	Username     string     `json:"username"`     // 用户名称
	Icon         string     `json:"icon"`         // 用户头像
	LoginAddress string     `json:"loginAddress"` // 用户登录地址
	CreateTime   util.HTime `json:"createTime"`   // 创建时间
}
