package image_api

import (
	"fmt"

	"github.com/2329497804/blogX-server/common"
	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models"
	"github.com/2329497804/blogX-server/service/log_service"
	"github.com/gin-gonic/gin"
)

type ImageApi struct {
}

type ImageListResponse struct {
	models.ImageModel
	WebPath string `json:"webPath"`
}

func (ImageApi) ImageListView(c *gin.Context) {
	var req common.PageInfo
	if err := c.ShouldBindQuery(&req); err != nil {
		res.FailWithError(err, c)
		return
	}
	_list, count, err := common.ListQuery[models.ImageModel](models.ImageModel{}, common.Options{
		PageInfo: req,
		Likes:    []string{"filename"},
	})
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var list = make([]ImageListResponse, 0)
	for _, v := range _list {
		list = append(list, ImageListResponse{
			ImageModel: v,
			WebPath:    v.WebPath(),
		})
	}
	res.OKWithList(list, count, c)
}

func (ImageApi) ImageDeleteView(c *gin.Context) {
	var req models.RemoveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithError(err, c)
		return
	}
	// 记录操作日志
	log := log_service.GetLog(c)
	log.ShowRequest()
	log.ShowResponse()

	var imageList []models.ImageModel
	global.DB.Find(&imageList, "id in ?", req.IDList)
	var successCount, failCount int64
	if len(imageList) > 0 {
		successCount = global.DB.Delete(&imageList).RowsAffected
	}
	failCount = int64(len(imageList)) - successCount

	msg := fmt.Sprintf("操作成功, 删除%d条, 失败%d条", successCount, failCount)

	res.OKWithMsg(msg, c)
}
