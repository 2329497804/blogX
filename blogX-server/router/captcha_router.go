package router

import (
	"github.com/2329497804/blogX-server/api"
	"github.com/gin-gonic/gin"
)

func CaptchaRouter(r *gin.RouterGroup) {
	r = r.Group("/captcha")

	app := api.App.CaptchaApi
	r.GET("", app.CaptChaView)
}
