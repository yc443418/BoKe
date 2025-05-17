// jwt
// @author chen

package core

import (
	"errors"
	"fmt"
	"gin-admin-api/config"
	"gin-admin-api/constant"
	"gin-admin-api/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type userStdClaims struct {
	model.JwtAdmin
	jwt.StandardClaims
}

var (
	TokenExpireDuration = time.Duration(config.Config.Token.ExpireTime) * time.Hour // token过期时间
	Secret              = []byte(config.Config.Token.Secret)                        // token密钥
	Issuer              = config.Config.Token.Issuer                                // 签发人
	ErrAbsent           = "token absent"                                            // 令牌不存在
	ErrInvalid          = "token invalid"                                           // 令牌无效
)

// GenerateTokenByAdmin 根据用户信息生成token
func GenerateTokenByAdmin(admin model.SysAdmin) (string, error) {
	var jwtAdmin = model.JwtAdmin{
		ID:         admin.ID,
		Username:   admin.Username,
		Nickname:   admin.Nickname,
		Status:     admin.Status,
		Icon:       admin.Icon,
		Email:      admin.Email,
		Phone:      admin.Phone,
		Note:       admin.Note,
		ExpireTime: time.Now().Add(TokenExpireDuration).Unix(),
	}
	c := userStdClaims{
		jwtAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

// ValidateToken 解析token
func ValidateToken(tokenString string) (*model.JwtAdmin, error) {
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
	return &claims.JwtAdmin, err
}

// GetAdmin 返回admin信息
func GetAdmin(c *gin.Context) (*model.JwtAdmin, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return nil, errors.New("can't get api")
	}
	admin, ok := u.(*model.JwtAdmin)
	if ok {
		return admin, nil
	}
	return nil, errors.New("can't convert to api struct")
}
