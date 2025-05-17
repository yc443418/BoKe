// 状态码和状态信息定义
// @author chen

package result

// Codes 定义状态
type Codes struct {
	Message                 map[uint]string
	Success                 uint
	Failed                  uint
	NoAuth                  uint
	AuthFormatError         uint
	InvalidToken            uint
	EmailCodeError          uint
	MissParameter           uint
	UsernameAlreadyExists   uint
	EmailAlreadyExists      uint
	PasswordError           uint
	EmailCodeExpire         uint
	EmailCodeIsFalse        uint
	PasswordNotTrue         uint
	StatusIsEnable          uint
	UserIsNotExist          uint
	CodeHasExpired          uint
	CaptchaNotTrue          uint
	FileUploadError         uint
	ArticleIsExist          uint
	ArticleIsNotExist       uint
	ArticleIsNotYou         uint
	ReplyUserNotExist       uint
	BArticleCommentNotExist uint
	SecondCommentNotTop     uint
	NotYouComment           uint
}

// ApiCode 状态码
var ApiCode = &Codes{
	Success:                 200,
	Failed:                  501,
	NoAuth:                  401,
	AuthFormatError:         402,
	InvalidToken:            403,
	EmailCodeError:          600,
	MissParameter:           601,
	UsernameAlreadyExists:   602,
	EmailAlreadyExists:      603,
	PasswordError:           604,
	EmailCodeExpire:         605,
	EmailCodeIsFalse:        606,
	PasswordNotTrue:         607,
	StatusIsEnable:          608,
	UserIsNotExist:          609,
	CodeHasExpired:          610,
	CaptchaNotTrue:          611,
	FileUploadError:         612,
	ArticleIsExist:          613,
	ArticleIsNotExist:       614,
	ArticleIsNotYou:         615,
	ReplyUserNotExist:       616,
	BArticleCommentNotExist: 617,
	SecondCommentNotTop:     618,
	NotYouComment:           619,
}

// 状态信息初始化
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.Success:                 "成功",
		ApiCode.Failed:                  "失败",
		ApiCode.NoAuth:                  "无权限",
		ApiCode.AuthFormatError:         "请求头中的auth格式错误",
		ApiCode.InvalidToken:            "无效的token,请重新登录",
		ApiCode.EmailCodeError:          "生成邮箱验证码错误",
		ApiCode.MissParameter:           "缺少参数",
		ApiCode.UsernameAlreadyExists:   "账号已存在，请重新输入",
		ApiCode.EmailAlreadyExists:      "邮箱已存在，请重新输入",
		ApiCode.PasswordError:           "两次输入的密码不一致，请新输入",
		ApiCode.EmailCodeExpire:         "邮箱验证码过期",
		ApiCode.EmailCodeIsFalse:        "输入的邮箱验证码不正确，请重新输入",
		ApiCode.PasswordNotTrue:         "密码不正确，请重新输入",
		ApiCode.StatusIsEnable:          "您的邮箱已停用，请联系官方解封",
		ApiCode.UserIsNotExist:          "用户不存在，请重新输入",
		ApiCode.CodeHasExpired:          "验证码已过期。",
		ApiCode.CaptchaNotTrue:          "验证码不正确，请重新输入！",
		ApiCode.FileUploadError:         "文件上传错误",
		ApiCode.ArticleIsExist:          "文章标题已存在，请重新输入！",
		ApiCode.ArticleIsNotExist:       "文章不存在",
		ApiCode.ArticleIsNotYou:         "不是你的文章，不能操作",
		ApiCode.ReplyUserNotExist:       "回复人不存在",
		ApiCode.BArticleCommentNotExist: "评论不存在",
		ApiCode.SecondCommentNotTop:     "二级评论不能置顶操作",
		ApiCode.NotYouComment:           "不是自己的评论不能操作",
	}
}

// GetMessage 供外部调用
func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
