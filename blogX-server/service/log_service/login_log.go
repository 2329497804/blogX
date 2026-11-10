package log_service

import (
	"github.com/2329497804/blogX-server/core"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models"
	"github.com/2329497804/blogX-server/models/enum"
	"github.com/2329497804/blogX-server/utils/jwts"
	"github.com/gin-gonic/gin"
)

// NewLoginSuccess 登陆成功的日志
func NewLoginSuccess(c *gin.Context, loginType enum.LoginType) {
	ip := c.ClientIP()
	addr := core.GetIPAddr(ip)

	userID := uint(0)
	userName := ""
	claims, err := jwts.ParseTokenByGin(c)
	if err == nil && claims != nil {
		userID = claims.UserID
		userName = claims.Username
	}

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
