package config

const (
	SUCCEE                    = 200
	ERROR                     = 500
	ERROR_USERNAME_USED       = 400
	ERROR_PASSWORD_WRONG      = 401
	ERROR_PASSWORD_EMPYT      = 402
	ERROR_USER_NOT_EXIST      = 404
	ERROR_DATABASE_ERROR      = 501
	ERROR_ENITY_NOT_EXIST     = 405 // 传参不成功
	ERROR_TOKEN_NOT_EXIST     = 406
	ERROR_TOKEN_EXPIRED       = 407
	ERROR_TOKEN_WRONG         = 408 //无法生成token
	ERROR_EMPYT_FIELD         = 409
	ERROR_CREATE_RANDOM_ERROR = 410
	ERROR_QUERY_USER_ERROR    = 414
	ERROR_WEBSOCKET_UPGRADE   = 415
	ERROR_AUTH_FAILED         = 416
	ERROR_INVALID_DATA        = 417
	ERROR_PASSWORD_CONFIRM    = 418
	ERROR_PASSWORD_WEAK       = 419
)

var errMsg = map[int]string{
	SUCCEE:                    "操作成功",
	ERROR:                     "操作失败",
	ERROR_DATABASE_ERROR:      "数据库错误",
	ERROR_ENITY_NOT_EXIST:     "实体不存在",
	ERROR_USERNAME_USED:       "用户名已被使用",
	ERROR_PASSWORD_WRONG:      "密码错误",
	ERROR_PASSWORD_EMPYT:      "密码不能为空",
	ERROR_USER_NOT_EXIST:      "用户不存在",
	ERROR_TOKEN_NOT_EXIST:     "Token不存在",
	ERROR_TOKEN_EXPIRED:       "Token已过期", // 后续需要规定具体过期时间
	ERROR_TOKEN_WRONG:         "Token无效",
	ERROR_EMPYT_FIELD:         "存在未填写的字段",
	ERROR_CREATE_RANDOM_ERROR: "生成随机数失败", // 创建随机数时出错
	ERROR_QUERY_USER_ERROR:    "查询用户信息失败",
	ERROR_WEBSOCKET_UPGRADE:   "WebSocket升级失败",
	ERROR_AUTH_FAILED:         "认证失败",
	ERROR_INVALID_DATA:        "参数绑定失败",
	ERROR_PASSWORD_CONFIRM:    "密码未一致",
	ERROR_PASSWORD_WEAK:       "密码强度过低",
}

func GetErrMsg(code int) string {
	return errMsg[code]
}
