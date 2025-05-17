// 文章相关模型
// @author chen

package model

import util "gin-front-api/utils"

// BArticle 文章模型
type BArticle struct {
	ID              uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                        // ID
	UserId          uint       `gorm:"column:user_id;comment:'用户id';NOT NULL" json:"userId"`                        // 用户id
	CategoryId      uint       `gorm:"column:category_id;comment:'分类id';NOT NULL" json:"categoryId"`                // 分类id
	TagId           uint       `gorm:"column:tag_id;comment:'标签id';NOT NULL" json:"tagId"`                          // 标签id
	Title           string     `gorm:"column:title;varchar(150);comment:'标题'" json:"title"`                         // 文章标题
	Image           string     `gorm:"column:image;varchar(200);comment:'封面地址'" json:"image"`                       // 封面地址
	Summary         string     `gorm:"column:summary;varchar(255);comment:'文章简介'" json:"summary"`                   // 文章简介
	MarkdownContent string     `gorm:"column:markdown_content;comment:'文章内容md版'" json:"markdownContent"`            // 文章内容md版
	Content         string     `gorm:"column:content;comment:'文章内容'" json:"content"`                                // 文章内容
	TopType         uint       `gorm:"column:top_type;comment:'是否发布(1-未发布 2-已发布)';DEFAULT 1" json:"topType"`        // 是否置顶(1-未置顶 2-已置顶)
	Quantity        uint       `gorm:"column:quantity;default:0;comment:'文章阅读量';NOT NULL" json:"quantity"`          // 文章阅读量
	GoodCount       uint       `gorm:"column:good_count;default:0;comment:'文章点赞量';NOT NULL" json:"goodCount"`       // 文章点赞量
	CommentCount    uint       `gorm:"column:comment_count;default:0;comment:'文章评论量';NOT NULL" json:"commentCount"` // 文章评论量
	CreateTime      util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`                // 创建时间
	Status          uint       `gorm:"column:status; comment:'状态（1-已删除，2-未删除）';DEFAULT 2" json:"status"`            // 状态（1-已删除，2-未删除）
}

func (BArticle) TableName() string {
	return "b_article"
}

// BArticleDto 新增文章参数
type BArticleDto struct {
	Id              uint   `json:"id"`              // 文章id
	CategoryId      uint   `json:"categoryId"`      // 分类id
	TagId           uint   `json:"tagId"`           // 标签id
	Title           string `json:"title"`           // 文章标题
	Image           string `json:"image"`           // 封面地址
	Summary         string `json:"summary"`         // 文章简介
	MarkdownContent string `json:"markdownContent"` // 文章内容md版
	Content         string `json:"content"`         // 文章内容
}

// BArticleVo 文章视图列表对象
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
	Icon         string     `json:"icon"`         // 头像
	LoginAddress string     `json:"loginAddress"` // 登录地址
	Summary      string     `json:"summary"`      // 文章简介
	TopType      uint       `json:"topType"`      // 是否置顶(1-未置顶 2-已置顶)
	Quantity     uint       `json:"quantity"`     // 文章阅读量
	GoodCount    uint       `json:"goodCount"`    // 点赞量
	CommentCount uint       `json:"commentCount"` // 评论量
	CreateTime   util.HTime `json:"createTime"`   // 创建时间
	Status       uint       `json:"status"`       // 状态（1-已删除，2-未删除）
}

// BArticleDetailVo 文章详情
type BArticleDetailVo struct {
	ID              uint       `json:"id"`              // ID
	Image           string     `json:"image"`           // 封面地址
	Title           string     `json:"title"`           // 文章标题
	MarkdownContent string     `json:"markdownContent"` // 文章内容md版
	Content         string     `json:"content"`         // 文章内容
	CategoryId      uint       `json:"categoryId"`      // 分类id
	CategoryName    string     `json:"categoryName"`    // 分类名称
	TagId           uint       `json:"tagId"`           // 标签id
	TagName         string     `json:"tagName"`         // 标签名称
	UserId          uint       `json:"userId"`          // 用户id
	Username        string     `json:"username"`        // 用户名称
	Icon            string     `json:"icon"`            // 用户头像
	Summary         string     `json:"summary"`         // 文章简介
	Quantity        uint       `json:"quantity"`        // 文章阅读量
	GoodCount       uint       `json:"goodCount"`       // 文章点赞数量
	CommentCount    uint       `json:"commentCount"`    // 评论量
	HaveLike        bool       `json:"haveLike"`        // 是否点赞(false:否 true:是)
	CreateTime      util.HTime `json:"createTime"`      // 创建时间
}
