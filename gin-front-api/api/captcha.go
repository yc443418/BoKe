// 验证码相关接口
// @author chen

package api

import (
	"gin-front-api/component"
	"gin-front-api/global"
	"gin-front-api/result"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store = component.RedisStore{}

// Captcha 登录验证码
// @Summary 登录验证码
// @Tags 登录验证码相关接口
// @Produce json
// @Description 登录验证码
// @Success 200 {object} result.Result
// @Router /api/captcha [get]
func Captcha(c *gin.Context) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString
	// 配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          6,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	lid, lb64s, _ := captcha.Generate()
	global.Log.Infof("验证码为：%s", component.RedisStore{}.Get(lid, true))
	result.Success(c, map[string]any{"idKey": lid, "image": lb64s})
}
