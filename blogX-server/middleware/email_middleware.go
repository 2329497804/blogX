package middleware

import (
	"bytes"
	"io"

	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/utils/emails"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type EmailMiddlewareRequest struct {
	EmailID   string `json:"emailID" binding:"required"`
	EmailCode string `json:"emailCode" binding:"required"`
}

func EmailMiddleWare(c *gin.Context) {
	// 请求体阅后即焚
	body, err := c.GetRawData()
	if err != nil {
		logrus.Errorf("获取请求体失败: %s", err)
		res.FailWithMsg("邮箱验证失败", c)
		c.Abort()
		return
	}
	// 写回 body
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
	var req EmailMiddlewareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	// 验证邮箱验证码
	_info, ok := global.EmailStore.Load(req.EmailID)
	if !ok {
		res.FailWithMsg("邮箱验证失败", c)
		c.Abort()
		return
	}
	info, ok := _info.(emails.EmailStoreInfo)
	if !ok {
		res.FailWithMsg("邮箱验证失败", c)
		c.Abort()
		return
	}
	if req.EmailCode != info.Code {
		res.FailWithMsg("邮箱验证码错误", c)
		c.Abort()
		return
	}
	// email 写入上下文
	c.Set("email", info.Email)
	// 写回 body
	c.Request.Body = io.NopCloser(bytes.NewReader(body))

	c.Next()

}
