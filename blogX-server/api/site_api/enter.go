package site_api

import (
	"fmt"

	"github.com/2329497804/blogX-server/service/log_service"
	"github.com/gin-gonic/gin"
)

type SiteApi struct {
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	fmt.Println("site info")
	c.JSON(200, gin.H{"code": 0, "data": "站点信息"})
	return
}

func (SiteApi) SiteUpdateView(c *gin.Context) {
	log := log_service.NewActionLogByGin(c)

	log.Save()

	c.JSON(200, gin.H{"code": 0, "msg": "站点信息"})
	return
}
