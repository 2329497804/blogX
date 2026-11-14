package site_api

import (
	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models/enum"
	"github.com/2329497804/blogX-server/service/log_service"
	"github.com/2329497804/blogX-server/utils/jwts"
	"github.com/gin-gonic/gin"
)

type SiteApi struct {
}

type SiteInfoRequest struct {
	Name string `uri:"name"`
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	var req SiteInfoRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	// 站点信息不需要登录
	if req.Name == "site" {
		res.OKWithData(global.Config.Site, c)
		return
	}

	// 判断是否为管理员
	ok, err := isAdmin(c)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	if !ok {
		res.FailWithMsg("权限不足", c)
	}

	//敏感信息打码返回
	var data any
	switch req.Name {
	case "email":
		v := global.Config.Email
		v.AuthCode = "********"
		data = v
	case "qq":
		v := global.Config.QQ
		v.AppKey = "********"
		data = v
	case "qiNiu":
		v := global.Config.QiNiu
		v.SecretKey = "********"
		data = v
	case "ai":
		v := global.Config.AI
		v.Secret = "********"
		data = v
	default:
		res.FailWithMsg("不存在的配置", c)
		return
	}
	res.OKWithData(data, c)
	return
}

func isAdmin(c *gin.Context) (bool, error) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		return false, err
	}
	if claims.Role != enum.AdminRole {
		return false, nil
	}
	return true, nil
}

// SiteInfoQQView QQ登陆地址
func (SiteApi) SiteInfoQQView(c *gin.Context) {
	res.OKWithData(global.Config.QQ.Url(), c)
}

type SiteUpdateRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required" label:"年龄"`
}

func (SiteApi) SiteUpdateView(c *gin.Context) {
	log := log_service.GetLog(c)

	var req SiteUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.SetError("参数绑定失败", err)
		res.FailWithError(err, c)
		return
	}

	res.OKWithMsg("更新成功", c)
	return
}
