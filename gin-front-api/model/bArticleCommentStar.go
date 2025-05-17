// 文章评论点赞相关模型
// @author chen

package model

// BArticleCommentStar 文章评论点赞模型
type BArticleCommentStar struct {
	ID        uint `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`         // ID
	UserId    uint `gorm:"column:user_id;comment:'点赞用户id';NOT NULL" json:"userId"`       // 点赞用户id
	CommentId uint `gorm:"column:comment_id;comment:'文章评论id';NOT NULL" json:"commentId"` // 文章评论id
}

func (BArticleCommentStar) TableName() string {
	return "b_article_comment_star"
}
