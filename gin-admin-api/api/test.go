// 测试相关接口
// @author chen

package api

import (
	"gin-admin-api/result"
	"github.com/gin-gonic/gin"
)

// Success 成功测试
// @Summary 成功测试接口
// @Tags 测试相关接口
// @Produce json
// @Description 成功测试接口
// @Success 200 {object} result.Result
// @router /api/success [get]
// @Security ApiKeyAuth
func Success(c *gin.Context) {
	result.Success(c, 200)
}

// Failed 失败测试
// @Summary 失败测试接口
// @Tags 测试相关接口
// @Produce json
// @Description 失败测试接口
// @Success 200 {object} result.Result
// @router /api/failed [get]
// @Security ApiKeyAuth
func Failed(c *gin.Context) {
	result.Failed(c, int(result.ApiCode.Failed), result.ApiCode.GetMessage(result.ApiCode.Failed))
}
