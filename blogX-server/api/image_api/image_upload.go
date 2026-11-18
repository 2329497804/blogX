package image_api

import (
	"fmt"
	"io"

	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models"
	"github.com/2329497804/blogX-server/service/qiniu_service"
	"github.com/2329497804/blogX-server/utils/files"
	"github.com/2329497804/blogX-server/utils/hashs"
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
	fileName := fileHeader.Filename
	var suffix string
	suffix, err = files.ImageSuffixJudge(fileName)
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
	hash := hashs.MD5(byteData)
	// 判断是否存在
	var m models.ImageModel
	err = global.DB.Take(&m, "hash = ?", hash).Error
	if err == nil {
		// 文件已存在
		logrus.Infof("上传图片重复: %s <==> %s, Hash:%s", m.Filename, fileName, hash)
		res.OK(m.WebPath(), "图片上传成功", c)
		return
	}
	// 存储文件
	url, err := qiniu_service.SendReader(file, c)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	fileDir := global.Config.Upload.Image.Dir
	filePath := fmt.Sprintf("uploads/%s/%s.%s", fileDir, hash, suffix)
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

	res.OK(m.WebPath(), "图片上传成功: " + url, c)
}
