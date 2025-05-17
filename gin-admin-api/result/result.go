// 结构数据定义
// @author chen

package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Result 结构体
type Result struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据
}

// Success 成功
func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{}
	res.Code = int(ApiCode.Success)
	res.Message = ApiCode.GetMessage(ApiCode.Success)
	res.Data = data
	c.JSON(http.StatusOK, res)
}

// Failed 失败
func Failed(c *gin.Context, code int, message string) {
	res := Result{}
	res.Code = code
	res.Message = message
	res.Data = gin.H{}
	c.JSON(http.StatusOK, res)
}
