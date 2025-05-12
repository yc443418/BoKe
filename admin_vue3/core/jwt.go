// jwt鉴权
// @author chen

package core

import (
	"admin_vue3/config"
	"admin_vue3/model"
	"time"

	"github.com/dgrijalva/jwt-go"
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
