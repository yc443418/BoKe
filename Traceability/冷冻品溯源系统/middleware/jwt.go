package middleware

import (
	ErrMsg "Traceability/config"
	"Traceability/utils"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			tokenString string
			code        int
		)

		// 获取 Token
		tokenQuery := c.Query("token")
		if tokenQuery != "" {
			tokenString = tokenQuery
		} else {
			tokenString = c.Request.Header.Get("Authorization")
		}

		if tokenString == "" {
			code = ErrMsg.ERROR_TOKEN_NOT_EXIST
			utils.ResultReturnErr(c, tokenString, code)
			c.Abort()
			return
		}
		// 解析 Token
		if claims, err := utils.ParseToken(tokenString); err != nil {
			code = ErrMsg.ERROR_TOKEN_NOT_EXIST
			utils.ResultReturnErr(c, err, code)
			c.Abort()
			return
		} else {
			c.Set("claims", claims) // 将用户信息存储在上下文中
			c.Next()                // 继续处理请求
			return
		}
	}
}
