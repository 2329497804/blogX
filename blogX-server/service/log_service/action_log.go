package log_service

import (
	"github.com/2329497804/blogX-server/core"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models"
	"github.com/2329497804/blogX-server/models/enum"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ActionLog 操作日志
type ActionLog struct {
	c     *gin.Context
	level enum.LogLevelType
	title string
}

func NewActionLogByGin(c *gin.Context) *ActionLog {
	return &ActionLog{
		c: c,
	}
}

func (ac *ActionLog) SetLevel(l enum.LogLevelType) {
	ac.level = l
}

func (ac *ActionLog) SetTitle(title string) {
	ac.title = title
}

func (ac *ActionLog) Save() {
	ip := ac.c.ClientIP()
	addr := core.GetIPAddr(ip)
	userID := uint(1)
	
	err := global.DB.Create(&models.LogModel{
		LogType: enum.ActionLogType,
		Title:   ac.title,
		Content: "",
		Level:   ac.level,
		UserID:  userID,
		IP:      ip,
		Addr:    addr,
	}).Error
	if err != nil {
		logrus.Errorf("日志创建失败: %s", err)
	}
}
