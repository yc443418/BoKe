// 文章点赞相关模型
// @author chen

package model

// BArticleStar 文章点赞模型
type BArticleStar struct {
	ID        uint `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`       // ID
	UserId    uint `gorm:"column:user_id;comment:'点赞用户id';NOT NULL" json:"userId"`     // 点赞用户id
	ArticleId uint `gorm:"column:article_id;comment:'文章id';NOT NULL" json:"articleId"` // 文章id
}

func (BArticleStar) TableName() string {
	return "b_article_star"
}
