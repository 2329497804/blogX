package middleware

import (
	"bytes"
	"io"

	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CaptchaMiddlewareRequest struct {
	ID   string `json:"captchaID"`
	Code string `json:"captchaCode"`
}

func CaptchaMiddleware(c *gin.Context) {
	if !global.Config.Site.Login.Captcha {
		return
	}
	// 请求体阅后即焚
	body, err := c.GetRawData()
	if err != nil {
		logrus.Errorf("获取请求体失败: %s", err)
		res.FailWithMsg("图形验证失败", c)
		c.Abort()
		return
	}
	// 写回 body
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
	var req CaptchaMiddlewareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithMsg("图形验证失败", c)
		c.Abort()
		return
	}
	if !global.CaptchaStore.Verify(req.ID, req.Code, true) {
		res.FailWithMsg("验证码错误", c)
		c.Abort()
		return
	}
	// 写回 body
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
}
