package utils

import (
	ErrMsg "Traceability/config"
	"Traceability/global"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"strings"
	"time"
)

var mySignKey = []byte(global.CONFIG.Jwt.Secret)

// 获取 JWT 过期时间
func GetExpireTime() time.Duration {
	return time.Duration(global.CONFIG.Jwt.ExpireTime) * time.Hour
}

// 获取 JWT 生效时间
func GetNotBefore() time.Duration {
	return time.Duration(global.CONFIG.Jwt.NotBefore) * time.Second
}

// NeedClaims 是自定义的 JWT Claims，包含需要的字段
type NeedClaims struct {
	UserID   string `json:"user_id"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}

// 生成 JWT token
func GenRegisterToken(claims NeedClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	return token.SignedString(mySignKey)
}

// 解析 JWT token
func ParseToken(tokenString string) (*NeedClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &NeedClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySignKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("解析 token 错误: %w", err)
	}

	// 断言 token.Claims 为 *NeedClaims 类型
	if tokenstr, ok := token.Claims.(*NeedClaims); ok && token.Valid {
		return tokenstr, nil
	}

	return nil, errors.New("无效的 token")
}

func GetClaims(c *gin.Context) (*NeedClaims, error) {
	var err error
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {

		ResultReturnErr(c, "Token 为空", ErrMsg.ERROR_TOKEN_NOT_EXIST)
		return &NeedClaims{}, errors.New("token为空")
	}

	// 更灵活的Token提取方式
	var tokenString string
	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		tokenString = authHeader // 也允许不带Bearer前缀的情况
	}

	// 调试日志
	log.Printf("Attempting to parse token: %s", tokenString)

	// 解析Token
	claims, err := ParseToken(tokenString)
	return claims, err
}
