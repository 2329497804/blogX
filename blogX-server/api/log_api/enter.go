package log_api

import (
	"fmt"

	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models"
	"github.com/2329497804/blogX-server/models/enum"
	"github.com/gin-gonic/gin"
)

type LogApi struct {
}

type LogListRequest struct {
	Limit       int               `form:"limit"`
	Page        int               `form:"page"`
	Key         string            `form:"key"`
	LogType     enum.LogType      `form:"logType"`
	Level       enum.LogLevelType `form:"level"`
	UserID      uint              `form:"userID"`
	IP          string            `form:"ip"`
	LoginStatus bool              `form:"loginStatus"`
	ServiceName string            `form:"serviceName"`
}

type LogListResponse struct {
	models.LogModel
	UserNickname string `json:"userNickname"`
	UserAvatar   string `json:"userAvatar"`
}

func (LogApi) LogListView(c *gin.Context) {
	//分页、查询(精确查询、模糊查询)
	var req LogListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		res.FailWithError(err, c)
		return
	}

	var list []models.LogModel
	if req.Page > 20 {
		req.Page = 1
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit == 0 || req.Limit > 100 {
		req.Limit = 10
	}
	offset := (req.Page - 1) * req.Limit
	model := models.LogModel{
		LogType:     req.LogType,
		Level:       req.Level,
		UserID:      req.UserID,
		IP:          req.IP,
		LoginStatus: req.LoginStatus,
		ServiceName: req.ServiceName,
	}
	like := global.DB.Where("title like ?", fmt.Sprintf("%%%s%%", req.Key))
	global.DB.Preload("UserModel").Debug().Where(like).Where(model).Offset(offset).Limit(req.Limit).Find(&list)

	var count int64
	global.DB.Debug().Where(like).Where(model).Model(&models.LogModel{}).Count(&count)

	var _list []LogListResponse
	for _, v := range list {
		_list = append(_list, LogListResponse{
			LogModel:     v,
			UserNickname: v.UserModel.Nickname,
			UserAvatar:   v.UserModel.Avatar,
		})
	}

	res.OKWithList(_list, int(count), c)
}
