// 用户信息相关模型
// @author chen

package model

import util "gin-front-api/utils"

// UserMessage 消息模型
type UserMessage struct {
	ID             uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                                            // ID
	ReceivedUserId uint       `gorm:"column:received_user_id;comment:'接收人用户ID';DEFAULT NULL" json:"receivedUserId"`                    // 接收人用户ID
	ArticleId      uint       `gorm:"column:article_id;comment:'文章ID';DEFAULT NULL" json:"articleId"`                                  // 文章ID
	ArticleTitle   string     `gorm:"column:article_title;varchar(150);comment:'文章标题';DEFAULT NULL" json:"articleTitle"`               // 文章标题
	CommentId      uint       `gorm:"column:comment_id;comment:'评论ID';DEFAULT NULL" json:"commentId"`                                  // 评论ID
	SendUserId     uint       `gorm:"column:send_user_id;comment:'发送人用户ID';DEFAULT NULL" json:"sendUserId"`                            // 发送人用户ID
	SendUsername   string     `gorm:"column:send_username;varchar(150);comment:'发送人昵称';DEFAULT NULL" json:"sendUsername"`              // 发送人昵称
	SendUserIcon   string     `gorm:"column:send_user_icon;varchar(255);comment:'发送人头像';DEFAULT NULL" json:"sendUserIcon"`             // 发送人头像
	MessageType    int64      `gorm:"column:message_type;comment:'消息类型（1-回复我的，2-赞了文章，3-赞了评论，4-系统消息）';DEFAULT NULL" json:"messageType"` // 消息类型（1-回复我的，2-赞了文章，3-赞了评论，4-系统消息）
	MessageContent string     `gorm:"column:message_content;varchar(1000);comment:'消息内容';DEFAULT NULL" json:"messageContent"`          // 消息内容
	Status         uint       `gorm:"column:status;comment:'消息是否已读（1:未读，2:已读）';DEFAULT NULL" json:"status"`                            // 消息是否已读（1:未读，2:已读）
	CreateTime     util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`                                    // 创建时间
}

func (UserMessage) TableName() string {
	return "user_message"
}

// UserMessageVo 用户消息数据对象
type UserMessageVo struct {
	TotalCount       int64 `json:"totalCount"`       // 总数
	SystemCount      int64 `json:"systemCount"`      // 系统通知数
	ReplyCount       int64 `json:"replyCount"`       // 回复数
	StarArticleCount int64 `json:"starArticleCount"` // 点赞文章数
	StarCommentCount int64 `json:"starCommentCount"` // 点赞评论数
}
