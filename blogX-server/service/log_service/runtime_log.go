package log_service

import (
	"fmt"
	"strings"

	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models"
	"github.com/2329497804/blogX-server/models/enum"
	"github.com/sirupsen/logrus"
)

type RuntimeLog struct {
	level           enum.LogLevelType
	title           string
	log             *models.LogModel
	itemList        []string
	serviceName     string
	runtimeDateType enum.RuntimeDateType
}

func NewRuntimeLog(serviceName string, t enum.RuntimeDateType) *RuntimeLog {
	return &RuntimeLog{
		serviceName:     serviceName,
		runtimeDateType: t,
	}
}

func (r *RuntimeLog) Save() {
	//判断是创建还是更新
	var log models.LogModel

	global.DB.Find(&log, fmt.Sprintf("service_name = ? and log_type = ? and created_at >= DATE_SUB(NOW(), %s)",
		r.runtimeDateType.GetSQLTime()),
		r.serviceName,
		enum.RuntimeLogType,
	)

	content := strings.Join(r.itemList, "\n")

	if log.ID != 0 {
		//更新
		return
	}

	err := global.DB.Create(&models.LogModel{
		LogType: enum.RuntimeLogType,
		Title:   r.title,
		Level:   r.level,
		Content: content,
	}).Error
	if err != nil {
		logrus.Errorf("创建运行日志错误: %s", err)
		return
	}
}
