package log_service

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models"
	"github.com/2329497804/blogX-server/models/enum"
	"github.com/sirupsen/logrus"
)

type RuntimeLog struct {
	level           enum.LogLevelType
	title           string
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

func (r *RuntimeLog) SetLevel(l enum.LogLevelType) {
	r.level = l
}

func (r *RuntimeLog) SetTitle(title string) {
	r.title = title
}

func (r *RuntimeLog) SetLink(label, href string) {
	r.itemList = append(r.itemList, fmt.Sprintf("<div class=\"log_item link\"><div class=\"log_item_label\">%s</div><div class=\"log_item_content\"><a href=\"%s\" target=\"_blank\">%s</a></div></div>",
		label,
		href,
		href,
	))
}

func (r *RuntimeLog) SetImage(src string) {
	r.itemList = append(r.itemList, fmt.Sprintf("<div class=\"log_image\"><img src=\"%s\" alt=\"\"></div>",
		src,
	))
}

func (r *RuntimeLog) setItem(label string, value any, levelType enum.LogLevelType) {
	var v string
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice:
		byteData, _ := json.Marshal(value)
		v = string(byteData)
	default:
		v = fmt.Sprintf("%v", value)
	}

	r.itemList = append(r.itemList, fmt.Sprintf("<div class=\"log_item %s\"><div class=\"log_item_label\">%s</div><div class=\"log_item_content\">%s</div></div>",
		levelType,
		label,
		v,
	))
}
func (r *RuntimeLog) SetItem(label string, value any) {
	r.setItem(label, value, enum.InfoLogLevelType)
}
func (r *RuntimeLog) SetItemInfo(label string, value any) {
	r.setItem(label, value, enum.InfoLogLevelType)
}
func (r *RuntimeLog) SetItemWarn(label string, value any) {
	r.setItem(label, value, enum.WarnLogLevelType)
}
func (r *RuntimeLog) SetItemError(label string, value any) {
	r.setItem(label, value, enum.ErrorLogLevelType)
}

func (r *RuntimeLog) setNowTime() {
	r.itemList = append(r.itemList, fmt.Sprintf("<div class=\"log_time\">%s</div>",
		time.Now().Format("2006-01-02 15:04:05"),
	))
}

func (r *RuntimeLog) Save() {
	r.setNowTime()
	//判断是创建还是更新
	var log models.LogModel

	global.DB.Find(&log, fmt.Sprintf("service_name = ? and log_type = ? and created_at >= DATE_SUB(NOW(), %s)",
		r.runtimeDateType.GetSQLTime()),
		r.serviceName,
		enum.RuntimeLogType,
	)

	content := strings.Join(r.itemList, "\n")

	if log.ID != 0 {
		//已Save 更新
		c := strings.Join(r.itemList, "\n")
		newContent := log.Content + "\n" + c
		global.DB.Model(&log).Updates(map[string]any{
			"content": newContent,
		})
		r.itemList = []string{} // 清空老数据
		return
	}

	err := global.DB.Create(&models.LogModel{
		LogType:     enum.RuntimeLogType,
		Title:       r.title,
		Level:       r.level,
		Content:     content,
		ServiceName: r.serviceName,
	}).Error
	if err != nil {
		logrus.Errorf("创建运行日志错误: %s", err)
		return
	}
	r.itemList = []string{}
}
