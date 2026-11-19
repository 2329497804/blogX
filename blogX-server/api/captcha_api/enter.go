package captcha_api

import (
	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/global"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type CaptchaApi struct{}

type CaptChaResponse struct {
	ID      string `json:"id"`
	Captcha string `json:"captcha"`
}

func (CaptchaApi) CaptChaView(c *gin.Context) {
	// 验证码答案保存在内存中
	// 参数1：达到多少个验证码后触发一次过期清理
	// 参数2：验证码有效期
	//captchaStore = base64Captcha.NewMemoryStore(10240, 5*time.Minute)
	captchaStore := global.CaptchaStore

	// 数字验证码
	// height=80, width=240, length=5, maxSkew=0.7, dotCount=80
	captchaDriver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)

	captcha := base64Captcha.NewCaptcha(captchaDriver, captchaStore)
	// id、Base64 图片、答案、错误
	id, image, _, err := captcha.Generate()
	if err != nil {
		logrus.Error(err)
		res.FailWithMsg("生成验证码失败", c)
		return
	}
	v := CaptChaResponse{
		ID:      id,
		Captcha: image,
	}
	res.OKWithData(v, c)
}
