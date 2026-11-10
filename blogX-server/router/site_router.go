package router

import (
	"github.com/2329497804/blogX-server/api"
	"github.com/2329497804/blogX-server/middleware"
	"github.com/gin-gonic/gin"
)

func SiteRouter(r *gin.RouterGroup) {
	r = r.Group("/site")

	app := api.App.SiteApi
	r.GET("", app.SiteInfoView)
	r.PUT("", middleware.AdminMiddleware, app.SiteUpdateView)
}
