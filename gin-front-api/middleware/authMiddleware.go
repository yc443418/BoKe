// 鉴权中间件
// @author chen

package middleware

import (
	"gin-front-api/config"
	"gin-front-api/constant"
	"gin-front-api/core"
	"gin-front-api/result"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		// 请求头中的auth为空，无权限
		if authHeader == "" {
			result.Failed(c, int(result.ApiCode.NoAuth), result.ApiCode.GetMessage(result.ApiCode.NoAuth))
			c.Abort()
			return
		}
		// 请求头中的auth格式是否有错误
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == config.Config.Token.Headers) {
			result.Failed(c, int(result.ApiCode.AuthFormatError), result.ApiCode.GetMessage(result.ApiCode.AuthFormatError))
			c.Abort()
			return
		}
		// 无效的token,请重新登录
		mc, err := core.ValidateToken(parts[1])
		if err != nil {
			result.Failed(c, int(result.ApiCode.InvalidToken), result.ApiCode.GetMessage(result.ApiCode.InvalidToken))
			c.Abort()
			return
		} else {
			c.Set(constant.ContextKeyUserObj, mc)
			c.Next()
		}
	}
}
