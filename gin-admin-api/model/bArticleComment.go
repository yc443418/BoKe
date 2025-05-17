// 文章评论相关模型
// @author chen

package model

import util "gin-admin-api/utils"

// BArticleComment 文章评论
type BArticleComment struct {
	ID                uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                               // ID
	PCommentId        uint       `gorm:"column:p_comment_id;comment:'父级评论ID';DEFAULT NULL" json:"pCommentId"`                // 父级评论ID
	ArticleId         uint       `gorm:"column:article_id;comment:'文章id';NOT NULL" json:"articleId"`                         // 文章id
	Content           string     `gorm:"column:content;comment:'评论内容'" json:"content"`                                       // 评论内容
	UserId            uint       `gorm:"column:user_id;comment:'用户id';NOT NULL" json:"userId"`                               // 用户id
	Username          string     `gorm:"column:username;comment:'用户名称';DEFAULT NULL" json:"username"`                        // 用户名称
	Icon              string     `gorm:"column:icon;comment:'用户头像';DEFAULT NULL" json:"icon"`                                // 用户头像
	LoginAddress      string     `gorm:"column:login_address;comment:'用户登录地址';DEFAULT NULL" json:"loginAddress"`             // 用户登录地址
	ReplyUserId       uint       `gorm:"column:reply_user_id;comment:'回复人id';DEFAULT NULL" json:"replyUserId"`               // 回复人id
	ReplyUsername     string     `gorm:"column:reply_username;comment:'回复人用户名';DEFAULT NULL" json:"replyUsername"`           // 回复人用户名
	ReplyIcon         string     `gorm:"column:reply_icon;comment:'回复人用户头像';DEFAULT NULL" json:"replyIcon"`                  // 回复人用户头像
	ReplyLoginAddress string     `gorm:"column:reply_login_address;comment:'回复人登录地址';DEFAULT NULL" json:"replyLoginAddress"` // 回复人登录地址
	TopType           uint       `gorm:"column:top_type;default:1;comment:'是否置顶(1:未置顶 2:置顶)';NOT NULL" json:"topType"`       // 是否置顶(1:未置顶 2:置顶)
	CreateTime        util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`                       // 创建时间
	GoodCount         uint       `gorm:"column:good_count;comment:'点赞数量';NOT NULL" json:"goodCount"`                         // 点赞数量
	Status            uint       `gorm:"column:status;comment:'是否删除(1:已删除 2:未删除)';DEFAULT NULL" json:"status"`               // 是否删除(1:已删除 2:未删除)
	LikeType          uint       `gorm:"column:like_type;comment:'额外字段';DEFAULT NULL" json:"likeType"`                       // 额外字段 是否点赞(1:未点赞 2:已点赞)
}

func (BArticleComment) TableName() string {
	return "b_article_comment"
}

// BArticleCommentVo 文章评论对象
type BArticleCommentVo struct {
	ID           uint       `json:"id"`           // 评论ID
	Title        string     `json:"title"`        // 文章名称
	Content      string     `json:"content"`      // 评论内容
	Username     string     `json:"username"`     // 用户名称
	Icon         string     `json:"icon"`         // 用户头像
	LoginAddress string     `json:"loginAddress"` // 用户登录地址
	CreateTime   util.HTime `json:"createTime"`   // 创建时间
	Status       uint       `json:"status"`       // 是否审核(1:待审核 2:已审核)
}
