package image_api

import (
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models"
	"github.com/2329497804/blogX-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	var suffix string
	suffix, err = imageSuffixJudge(fileHeader.Filename)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	// 文件Hash
	file, err := fileHeader.Open()
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	defer file.Close()
	byteData, err := io.ReadAll(file)
	hash := utils.MD5(byteData)
	fileName := hash + "." + suffix
	// 判断是否存在
	var m models.ImageModel
	err = global.DB.Take(&m, "hash = ?", hash).Error
	if err == nil {
		// 文件已存在
		logrus.Infof("上传图片重复: %s <==> %s", m.Filename, fileName)
		res.OK(m.WebPath(), "图片上传成功", c)
		return
	}
	// 存储文件
	fileDir := global.Config.Upload.Image.Dir
	filePath := fmt.Sprintf("uploads/%s/%s", fileDir, fileName)
	err = c.SaveUploadedFile(fileHeader, filePath)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	m = models.ImageModel{
		Filename: fileName,
		Path:     filePath,
		Size:     fileSize,
		Hash:     hash,
	}
	err = global.DB.Create(&m).Error
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	res.OK(m.WebPath(), "图片上传成功", c)
}

func imageSuffixJudge(filename string) (suffix string, err error) {
	l := strings.Split(filename, ".")
	if len(l) < 2 {
		err = errors.New("错误的文件名")
		return
	}
	suffix = l[len(l)-1]
	whiteList := global.Config.Upload.Image.WhiteList
	ok := slices.Contains(whiteList, suffix)
	if !ok {
		err = errors.New("文件类型非法")
		return
	}
	return
}
