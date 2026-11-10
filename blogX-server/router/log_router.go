package router

import (
	"github.com/2329497804/blogX-server/api"
	"github.com/gin-gonic/gin"
)

func LogRouter(r *gin.RouterGroup) {
	app := api.App.LogApi
	r.GET("logs", app.LogListView)
}
