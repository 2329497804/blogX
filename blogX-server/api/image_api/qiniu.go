package image_api

import (
	"fmt"
	"uuid"

	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/service/qiniu_service"
	"github.com/gin-gonic/gin"
)

type QiNiuGenTokenResponse struct {
	Token  string `json:"token"`
	Key    string `json:"key"`
	Region string `json:"region"`
	URL    string `json:"url"`
	Size   int    `json:"size"`
}

func (ImageApi) QiNiuGenToken(c *gin.Context) {
	q := global.Config.QiNiu
	i := global.Config.Upload.Image
	if !q.Enable {
		res.FailWithMsg("未启用七牛云配置", c)
		return
	}
	token, err := qiniu_service.GenerateToken(c)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	uid := uuid.New().String()
	key := fmt.Sprintf("%s/%s/%s.png", q.Prefix, i.Dir, uid)
	url := fmt.Sprintf("%s/%s", q.URI, key)
	res.OKWithData(QiNiuGenTokenResponse{
		Token:  token,
		Key:    key,
		Region: q.Region,
		URL:    url,
		Size:   q.Size,
	}, c)
}
