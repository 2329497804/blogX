package router

import (
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/middleware"
	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(global.Config.Server.GinMode)

	r := gin.Default()

	r.Static("/static", "./uploads")

	nr := r.Group("/api")

	nr.Use(middleware.LogMiddleWare)
	SiteRouter(nr)
	LogRouter(nr)

	addr := global.Config.Server.Addr()
	if err := r.Run(addr); err != nil {
		panic(err)
	}
}
