package log_service

import (
	"fmt"

	"github.com/2329497804/blogX-server/core"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models"
	"github.com/2329497804/blogX-server/models/enum"
	"github.com/gin-gonic/gin"
)

// NewLoginSuccess 登陆成功的日志
func NewLoginSuccess(c *gin.Context, loginType enum.LoginType) {
	ip := c.ClientIP()
	addr := core.GetIPAddr(ip)

	token := c.GetHeader("token")
	fmt.Println(token)
	// TODO: 通过JWT获取用户id
	userID := uint(1)
	userName := ""

	global.DB.Create(&models.LogModel{
		LogType:     enum.LoginLogType,
		Title:       "用户登录",
		Content:     "",
		UserID:      userID,
		IP:          ip,
		Addr:        addr,
		LoginStatus: true,
		Username:    userName,
		Pwd:         "",
		LoginType:   loginType,
	})
}

// NewLoginFail 登陆失败的日志
func NewLoginFail(c *gin.Context, loginType enum.LoginType, msg, userName, pwd string) {
	ip := c.ClientIP()
	addr := core.GetIPAddr(ip)

	global.DB.Create(&models.LogModel{
		LogType:     enum.LoginLogType,
		Title:       "用户登录失败",
		Content:     msg,
		IP:          ip,
		Addr:        addr,
		LoginStatus: false,
		Username:    userName,
		Pwd:         pwd,
		LoginType:   loginType,
	})
}
