package site_api

import (
	"fmt"

	"github.com/2329497804/blogX-server/service/log_service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SiteApi struct {
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	fmt.Println("site info")
	c.JSON(200, gin.H{"code": 0, "data": "站点信息"})
	return
}

type SiteUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (SiteApi) SiteUpdateView(c *gin.Context) {
	log := log_service.GetLog(c)
	log.ShowRequest()
	log.ShowResponse()

	var req SiteUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf(err.Error())
	}
	fmt.Println("req:", req)

	c.JSON(200, gin.H{"code": 0, "msg": "站点信息"})
	return
}
