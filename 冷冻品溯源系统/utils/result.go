//package utils
//
//import (
//	ErrMsg "Traceability/config"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//// 成功返回的函数
//func ResultReturnOK(c *gin.Context, data interface{}, code int) {
//	msg := ErrMsg.GetErrMsg(code)
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  msg,
//		"data": data,
//	})
//}
//
//// 错误返回格式
//func ResultReturnErr(c *gin.Context, data interface{}, code int) {
//	msg := ErrMsg.GetErrMsg(code)
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  msg,
//		"data": data,
//	})
//}
//func TwoResultReturn(c *gin.Context, data1 interface{}, data2 interface{}, code int) {
//	msg := ErrMsg.GetErrMsg(code)
//	c.JSON(http.StatusOK, gin.H{
//		"code":  code,
//		"msg":   msg,
//		"data":  data1,
//		"data2": data2,
//	})
//}

package utils

import (
	ErrMsg "Traceability/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Result 定义了一个通用的响应结构体
type Result struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据
}

// 成功返回的函数
func ResultReturnOK(c *gin.Context, data interface{}, code int) {
	msg := ErrMsg.GetErrMsg(code)
	result := Result{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	c.JSON(http.StatusOK, result)
}

// 错误返回格式
func ResultReturnErr(c *gin.Context, data interface{}, code int) {
	msg := ErrMsg.GetErrMsg(code)
	result := Result{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	c.JSON(http.StatusOK, result)
}

// 返回两个数据的函数
func TwoResultReturn(c *gin.Context, data1 interface{}, data2 interface{}, code int) {
	msg := ErrMsg.GetErrMsg(code)
	result := Result{
		Code:    code,
		Message: msg,
		Data: gin.H{
			"data1": data1,
			"data2": data2,
		},
	}
	c.JSON(http.StatusOK, result)
}
