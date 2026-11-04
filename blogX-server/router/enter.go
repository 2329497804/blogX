package router

import (
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/middleware"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	nr := r.Group("/api")

	nr.Use(middleware.LogMiddleWare)
	SiteRouter(nr)

	addr := global.Config.Server.Addr()
	r.Run(addr)
}
