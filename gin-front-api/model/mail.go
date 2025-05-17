// 邮箱相关模型
// @author chen

package model

// MailDto 邮箱验证码
type MailDto struct {
	Email string `json:"email" validate:"required"` // 邮箱
}
