package banner_api

import (
	"fmt"

	"github.com/2329497804/blogX-server/common"
	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models"
	"github.com/gin-gonic/gin"
)

type BannerApi struct{}

type CreateBannerRequest struct {
	Cover string `json:"cover" binding:"required"`
	Href  string `json:"href"`
	Show  bool   `json:"show"`
}

func (BannerApi) CreateBannerView(c *gin.Context) {
	var req CreateBannerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithError(err, c)
		return
	}
	if err := global.DB.Create(&models.BannerModel{
		Cover: req.Cover,
		Href:  req.Href,
		Show:  req.Show,
	}).Error; err != nil {
		res.FailWithError(err, c)
		return
	}
	res.OKWithMsg("添加banner成功", c)
}

type ListBannerRequest struct {
	common.PageInfo
	Show bool `form:"show"`
}

func (BannerApi) ListBannerView(c *gin.Context) {
	var req ListBannerRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		res.FailWithError(err, c)
		return
	}
	list, count, err := common.ListQuery(models.BannerModel{
		Show: req.Show,
	}, common.Options{
		PageInfo: req.PageInfo,
	})
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	res.OKWithList(list, count, c)
}

func (BannerApi) RemoveBannerView(c *gin.Context) {
	var req models.RemoveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithError(err, c)
		return
	}
	var list []models.BannerModel
	global.DB.Find(&list, "id in ?", req.IDList)

	if len(list) > 0 {
		global.DB.Delete(&list)
	}
	successCount := len(list)
	failCount := len(req.IDList) - successCount

	msg := fmt.Sprintf("操作成功, 删除%d条, 失败%d条", successCount, failCount)
	res.OKWithMsg(msg, c)
}

func (BannerApi) UpdateBannerView(c *gin.Context) {
	var id models.IDRequest
	if err := c.ShouldBindUri(&id); err != nil {
		res.FailWithError(err, c)
		return
	}
	var req CreateBannerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithError(err, c)
		return
	}
	var m models.BannerModel
	if err := global.DB.Take(&m, id.ID).Error; err != nil {
		res.FailWithError(err, c)
		return
	}

	if err := global.DB.Model(&m).Updates(map[string]any{
		"cover": req.Cover,
		"href":  req.Href,
		"show":  req.Show,
	}).Error; err != nil {
		res.FailWithError(err, c)
		return
	}
	res.OKWithMsg("banner更新成功", c)
}
