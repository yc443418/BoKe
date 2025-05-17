// jwt工具类
// @author chen

package core

import (
	"errors"
	"fmt"
	"gin-front-api/config"
	"gin-front-api/constant"
	"gin-front-api/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type userStdClaims struct {
	model.JwtUser
	jwt.StandardClaims
}

var (
	TokenExpireDuration = time.Duration(config.Config.Token.ExpireTime) * time.Hour // token过期时间
	Secret              = []byte(config.Config.Token.Secret)                        // token密钥
	Issuer              = config.Config.Token.Issuer                                // 签发人
	ErrAbsent           = "token absent"                                            // 令牌不存在
	ErrInvalid          = "token invalid"                                           //令牌无效
)

// GenerateTokenByUser 根据用户信息生成token
func GenerateTokenByUser(user model.User) (string, error) {
	var jwtUser = model.JwtUser{
		ID:           user.ID,
		Username:     user.Username,
		Icon:         user.Icon,
		Email:        user.Email,
		Note:         user.Note,
		ExpireTime:   time.Now().Add(TokenExpireDuration).Unix(),
		LoginAddress: user.LoginAddress,
	}
	c := userStdClaims{
		jwtUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //过期时间
			Issuer:    Issuer,                                     //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

// ValidateToken 解析JWT
func ValidateToken(tokenString string) (*model.JwtUser, error) {
	if tokenString == "" {
		return nil, errors.New(ErrAbsent)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if token == nil {
		return nil, errors.New(ErrInvalid)
	}
	claims := userStdClaims{}
	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	return &claims.JwtUser, err
}

// GetUser 返回user信息
func GetUser(c *gin.Context) (*model.JwtUser, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return nil, errors.New("can't get api")
	}
	user, ok := u.(*model.JwtUser)
	if ok {
		return user, nil
	}
	return nil, errors.New("can't convert to api struct")
}
