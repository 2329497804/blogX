package router

import (
	"github.com/2329497804/blogX-server/api"
	"github.com/2329497804/blogX-server/middleware"
	"github.com/gin-gonic/gin"
)

func ImageRouter(r *gin.RouterGroup) {
	r = r.Group("/images")
	r.Use(middleware.AuthMiddleware)

	app := api.App.ImageApi
	r.POST("", app.ImageUploadView)
	r.GET("", middleware.AdminMiddleware, app.ImageListView)
	r.DELETE("", middleware.AdminMiddleware, app.ImageDeleteView)
}
