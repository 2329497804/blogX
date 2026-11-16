package site_api

import (
	"errors"
	"fmt"
	"os"

	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/conf"
	"github.com/2329497804/blogX-server/core"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models/enum"
	"github.com/2329497804/blogX-server/utils/jwts"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SiteApi struct {
}

type SiteInfoRequest struct {
	Name string `uri:"name"`
}

const code = "******"

func (SiteApi) SiteInfoView(c *gin.Context) {
	var req SiteInfoRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	// 站点信息不需要登录
	if req.Name == "site" {
		global.Config.Site.Version = global.Version
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
		v.AuthCode = code
		data = v
	case "qq":
		v := global.Config.QQ
		v.AppKey = code
		data = v
	case "qiNiu":
		v := global.Config.QiNiu
		v.SecretKey = code
		data = v
	case "ai":
		v := global.Config.AI
		v.Secret = code
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
	Name string `uri:"name" binding:"required"`
}

func (SiteApi) SiteUpdateView(c *gin.Context) {
	var req SiteUpdateRequest
	if err := c.ShouldBindUri(&req); err != nil {
		res.FailWithError(err, c)
		return
	}

	var err error
	var data any
	switch req.Name {
	case "site":
		var d conf.Site
		err = c.ShouldBindJSON(&d)
		data = d
	case "email":
		var d conf.Email
		err = c.ShouldBindJSON(&d)
		data = d
	case "qq":
		var d conf.QQ
		err = c.ShouldBindJSON(&d)
		data = d
	case "qiNiu":
		var d conf.QiNiu
		err = c.ShouldBindJSON(&d)
		data = d
	case "ai":
		var d conf.AI
		err = c.ShouldBindJSON(&d)
		data = d
	default:
		res.FailWithMsg("不存在的配置", c)
		return
	}
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	switch s := data.(type) {
	case conf.Site:
		err = updateSite(s)
		if err != nil {
			res.FailWithError(err, c)
			return
		}
		global.Config.Site = s
	case conf.Email:
		if s.AuthCode == code {
			s.AuthCode = global.Config.Email.AuthCode
		}
		global.Config.Email = s
	case conf.QQ:
		if s.AppKey == code {
			s.AppKey = global.Config.QQ.AppKey
		}
		global.Config.QQ = s
	case conf.QiNiu:
		if s.SecretKey == code {
			s.SecretKey = global.Config.QiNiu.SecretKey
		}
		global.Config.QiNiu = s
	case conf.AI:
		if s.Secret == code {
			s.Secret = global.Config.AI.Secret
		}
		global.Config.AI = s
	}
	// 持久化 修改配置文件
	if err = core.SetConfig(); err != nil {
		logrus.Error(err)
		res.FailWithMsg("修改配置文件失败"+err.Error(), c)
		return
	}

	res.OKWithMsg(fmt.Sprintf("更新站点%s配置成功", req.Name), c)
	return
}

func updateSite(site conf.Site) error {
	if site.Project.Icon == "" && site.Project.Title == "" && site.Project.WebPath == "" && site.Seo.Keywords == "" && site.Seo.Description == "" {
		return nil
	}
	if site.Project.WebPath == "" {
		return errors.New("请配置前端地址")
	}
	file, err := os.Open(site.Project.WebPath)
	if err != nil {
		return errors.New(site.Project.WebPath + " 文件不存在")
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		logrus.Errorf("goquery 解析失败 %s", err)
		return errors.New("文件解析失败")
	}

	if site.Project.Title != "" {
		doc.Find("title").SetText(site.Project.Title)
	}
	if site.Project.Icon != "" {
		selection := doc.Find("link[rel=\"icon\"]")
		if selection.Length() > 0 {
			// 有则修改
			selection.SetAttr("href", site.Project.Icon)
		} else {
			// 没有就创建
			doc.Find("head").AppendHtml(fmt.Sprintf("<link rel=\"icon\" href=\"%s\">", site.Project.Icon))
		}
	}
	if site.Seo.Keywords != "" {
		selection := doc.Find("meta[name=\"keywords\"]")
		if selection.Length() > 0 {
			// 有则修改
			selection.SetAttr("content", site.Seo.Keywords)
		} else {
			// 没有就创建
			doc.Find("head").AppendHtml(fmt.Sprintf("<meta name=\"keywords\" content=\"%s\">", site.Seo.Keywords))
		}
	}
	if site.Seo.Description != "" {
		selection := doc.Find("meta[name=\"description\"]")
		if selection.Length() > 0 {
			selection.SetAttr("content", site.Seo.Description)
		} else {
			doc.Find("head").AppendHtml(fmt.Sprintf("<meta name=\"description\" content=\"%s\">", site.Seo.Description))
		}
	}

	html, err := doc.Html()
	if err != nil {
		logrus.Errorf("生成html失败 %s", err)
		return errors.New("生成html失败")
	}

	err = os.WriteFile(site.Project.WebPath, []byte(html), 0644)
	if err != nil {
		logrus.Errorf("文件写入失败 %s", err)
		return errors.New("文件写入失败")
	}
	return nil
}
