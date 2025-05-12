package middleware

import (
	"Traceability/config"
	"Traceability/utils"
	"github.com/gin-gonic/gin"
)

// AdminOnly 管理员权限校验中间件
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			utils.ResultReturnErr(c, "未授权访问", config.ERROR_AUTH_FAILED)
			c.Abort()
			return
		}

		// 类型断言获取用户类型
		userClaims, ok := claims.(*utils.NeedClaims)
		if !ok || userClaims.UserType != "admin" {
			utils.ResultReturnErr(c, "需要管理员权限", config.ERROR_AUTH_FAILED)
			c.Abort()
			return
		}
		c.Next()
	}
}
