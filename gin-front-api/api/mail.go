// 邮箱相关api
// @author chen

package api

import (
	"fmt"
	"gin-front-api/config"
	"gin-front-api/constant"
	"gin-front-api/core"
	. "gin-front-api/core"
	"gin-front-api/global"
	"gin-front-api/model"
	"gin-front-api/result"
	"github.com/gin-gonic/gin"
	mailer "gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

// SendMail 邮箱注册验证码
// @Summary 邮箱注册验证码
// @Tags 邮箱注册相关接口
// @Produce json
// @Description 邮箱注册验证码
// @Param data body model.MailDto true "data"
// @Success 200 {object} result.Result
// @Router /api/qqMail [post]
func SendMail(c *gin.Context) {
	// 发送邮箱业务
	var dto model.MailDto
	_ = c.BindJSON(&dto)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	emailCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	core.RedisDb.Set(global.Ctx, constant.EmailCode+dto.Email, emailCode, time.Minute*5)
	now := time.Now()
	t := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	html := fmt.Sprintf(`<div>
		<div>
			尊敬的%s，您好！
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p>您于 %s 提交的邮箱验证，本次验证码为 %s，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
		</div>
		<div>
			<p>此邮箱为系统邮箱，请勿回复。</p>
		</div>	
	</div>`, dto.Email, t, emailCode)
	// 组装数据
	subjectRedis, err := core.RedisDb.Get(global.Ctx, constant.EmailSubjectKey).Result()
	if err != nil {
		var sysConfig model.SysConfig
		Db.Table("sys_config").Where("config_key = ?", constant.EmailSubjectKey).Scan(&sysConfig)
		core.RedisDb.Set(global.Ctx, sysConfig.ConfigKey, sysConfig.ConfigValue, 0)
		global.Log.Infof("从mysql中获取的邮箱主题: %s", sysConfig.ConfigValue)
		msg := mailer.NewMessage()
		msg.SetHeader("From", config.Config.Mail.Username)
		msg.SetHeader("To", dto.Email)
		msg.SetHeader("Subject", sysConfig.ConfigValue)
		msg.SetBody("text/html", html)
		dialer := mailer.NewDialer(
			config.Config.Mail.Host,
			config.Config.Mail.Port,
			config.Config.Mail.Username,
			config.Config.Mail.Password,
		)
		if err := dialer.DialAndSend(msg); err != nil {
			result.Failed(c, int(result.ApiCode.EmailCodeError), result.ApiCode.GetMessage(result.ApiCode.EmailCodeError))
			return
		}
		global.Log.Infof("发送者为: %s", config.Config.Mail.Username)
		global.Log.Infof("接受者为: %s", dto.Email)
		global.Log.Infof("发送主题为: %s", sysConfig.ConfigValue)
		global.Log.Infof("发送内容为: %s", html)
		global.Log.Infof("验证码为: %s", emailCode)
		result.Success(c, "true")
	} else {
		global.Log.Infof("从redis中获取的邮箱主题: %s", subjectRedis)
		msg := mailer.NewMessage()
		msg.SetHeader("From", config.Config.Mail.Username)
		msg.SetHeader("To", dto.Email)
		msg.SetHeader("Subject", subjectRedis)
		msg.SetBody("text/html", html)
		dialer := mailer.NewDialer(
			config.Config.Mail.Host,
			config.Config.Mail.Port,
			config.Config.Mail.Username,
			config.Config.Mail.Password,
		)
		if err := dialer.DialAndSend(msg); err != nil {
			result.Failed(c, int(result.ApiCode.EmailCodeError), result.ApiCode.GetMessage(result.ApiCode.EmailCodeError))
			return
		}
		global.Log.Infof("发送者为: %s", config.Config.Mail.Username)
		global.Log.Infof("接受者为: %s", dto.Email)
		global.Log.Infof("发送主题为: %s", subjectRedis)
		global.Log.Infof("发送内容为: %s", html)
		global.Log.Infof("验证码为: %s", emailCode)
		result.Success(c, "true")
	}
}
