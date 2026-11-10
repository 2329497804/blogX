package log_api

import (
	"github.com/2329497804/blogX-server/common"
	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/models"
	"github.com/2329497804/blogX-server/models/enum"
	"github.com/gin-gonic/gin"
)

type LogApi struct {
}

type LogListRequest struct {
	common.PageInfo
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
	list, count, err := common.ListQuery(models.LogModel{
		LogType:     req.LogType,
		Level:       req.Level,
		UserID:      req.UserID,
		IP:          req.IP,
		LoginStatus: req.LoginStatus,
		ServiceName: req.ServiceName,
	}, common.Options{
		PageInfo:     req.PageInfo,
		Likes:        []string{"title"},
		Preload:      []string{"UserModel"},
		Debug:        true,
		DefaultOrder: "id desc",
	})

	if err != nil {
		res.FailWithMsg(err.Error(), c)
	}
	var _list []LogListResponse
	for _, v := range list {
		_list = append(_list, LogListResponse{
			LogModel:     v,
			UserNickname: v.UserModel.Nickname,
			UserAvatar:   v.UserModel.Avatar,
		})
	}

	res.OKWithList(_list, count, c)
}
