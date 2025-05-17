// 文件相关接口
// @author chen

package api

import (
	"fmt"
	"gin-front-api/config"
	"gin-front-api/global"
	"gin-front-api/result"
	util "gin-front-api/utils"
	"github.com/gin-gonic/gin"
	"path"
	"strconv"
	"time"
)

// Upload 单上传
// @Summary 单上传
// @Tags 上传相关接口
// @Description 单上传
// @Produce json
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {object} result.Result
// @Router /api/upload [post]
// @Security ApiKeyAuth
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		result.Failed(c, int(result.ApiCode.FileUploadError), result.ApiCode.GetMessage(result.ApiCode.FileUploadError))
		return
	}
	now := time.Now()
	ext := path.Ext(file.Filename)
	fileName := strconv.Itoa(now.Nanosecond()) + ext
	filePath := fmt.Sprintf("%s%s%s%s",
		config.Config.UploadSettings.UploadDir,
		fmt.Sprintf("%04d", now.Year()),
		fmt.Sprintf("%02d", now.Month()),
		fmt.Sprintf("%04d", now.Day()))
	util.CreateDir(filePath)
	fullPath := filePath + "/" + fileName
	c.SaveUploadedFile(file, fullPath)
	global.Log.Infof("上传图片成功，地址为：%s", config.Config.UploadSettings.UploadHost+fullPath)
	result.Success(c, fullPath)
}
