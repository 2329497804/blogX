package router

import (
	"github.com/2329497804/blogX-server/api"
	"github.com/2329497804/blogX-server/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	r = r.Group("/user")

	app := api.App.UserApi

	r.POST("/login", middleware.CaptchaMiddleware, app.PWDLoginApi)
	r.GET("/detail", middleware.AuthMiddleware, app.UserDetailView)
	r.GET("/base", app.UserBaseInfoView)

	er := r.Group("/email")
	er.POST("/send", middleware.CaptchaMiddleware, app.SendEmailView)
	er.POST("", middleware.EmailMiddleWare, app.RegisterUserEmailView)
}
