package image_api

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/global"
	"github.com/gin-gonic/gin"
)

const (
	MB = 1024 * 1024
)

func (ImageApi) ImageUploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	// 文件大小判断
	fileSize := global.Config.Upload.Image.Size
	if fileHeader.Size > fileSize*MB {
		res.FailWithMsg(fmt.Sprintf("文件大小超过%dMB", fileSize), c)
		return
	}
	// 文件后缀判断
	if err = imageSuffixJudge(fileHeader.Filename); err != nil {
		res.FailWithError(err, c)
		return
	}
	// 存储文件
	fileDir := global.Config.Upload.Image.Dir
	fileName := fileHeader.Filename
	filePath := fmt.Sprintf("uploads/%s/%s", fileDir, fileName)
	err = c.SaveUploadedFile(fileHeader, filePath)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	res.OK("/"+filePath, "图片上传成功", c)
}

func imageSuffixJudge(filename string) (err error) {
	l := strings.Split(filename, ".")
	suffix := l[len(l)-1]
	whiteList := global.Config.Upload.Image.WhiteList
	ok := slices.Contains(whiteList, suffix)
	if !ok {
		err = errors.New("文件类型非法")
	}
	return
}
