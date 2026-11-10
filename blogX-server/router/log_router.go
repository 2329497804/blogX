package router

import (
	"github.com/2329497804/blogX-server/api"
	"github.com/2329497804/blogX-server/middleware"
	"github.com/gin-gonic/gin"
)

func LogRouter(r *gin.RouterGroup) {
	r = r.Group("/logs")
	r.Use(middleware.AdminMiddleware)

	app := api.App.LogApi
	r.GET("", app.LogListView)
	r.GET("/:id", app.LogReadView)
	r.DELETE("", app.LogRemoveView)
}
