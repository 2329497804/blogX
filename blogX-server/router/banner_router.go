package router

import (
	"github.com/2329497804/blogX-server/api"
	"github.com/2329497804/blogX-server/middleware"
	"github.com/gin-gonic/gin"
)

func BannerRouter(r *gin.RouterGroup) {
	r = r.Group("/banner")

	app := api.App.BannerApi
	r.GET("", app.ListBannerView)
	r.POST("", middleware.AdminMiddleware, app.CreateBannerView)
	r.PUT("/:id", middleware.AdminMiddleware, app.UpdateBannerView)
	r.DELETE("", middleware.AdminMiddleware, app.RemoveBannerView)
}
