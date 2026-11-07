package site_api

import (
	"fmt"
	"time"

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

type SiteUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (SiteApi) SiteUpdateView(c *gin.Context) {
	log := log_service.GetLog(c)
	log.ShowRequestHeader()
	log.ShowRequest()
	log.ShowResponseHeader()
	log.ShowResponse()
	log.SetTitle("更新站点")
	log.SetItem("请求时间", time.Now())
	log.SetImage("https://cdn-test.zjzw.com.cn/course/cover/20261105/c3d8425f-8a76-4b3d-b468-7849bbaefde7.png")
	log.SetLink("go学习", "https://www.fengfengzhidao.com")

	var req SiteUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.SetError("参数绑定失败", err)
	}

	log.SetItem("结构体", req)
	log.Save()
	log.SetItem("切片", []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"})
	log.Save()
	log.SetItem("数字", 123)
	log.Save()

	c.JSON(200, gin.H{"code": 0, "msg": "站点信息"})
	return
}
