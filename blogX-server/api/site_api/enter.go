package site_api

import (
	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/service/log_service"
	"github.com/gin-gonic/gin"
)

type SiteApi struct {
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	res.OKWithData("站点信息", c)
	return
}

type SiteUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (SiteApi) SiteUpdateView(c *gin.Context) {
	log := log_service.GetLog(c)

	var req SiteUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.SetError("参数绑定失败", err)
		res.FailWithMsg(err.Error(), c)
		return
	}

	res.OKWithMsg("更新成功", c)
	return
}
