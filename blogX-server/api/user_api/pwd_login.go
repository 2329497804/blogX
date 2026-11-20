package user_api

import (
	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models"
	"github.com/2329497804/blogX-server/utils/jwts"
	"github.com/2329497804/blogX-server/utils/pwd"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PWDLoginRequest struct {
	Val string `json:"val" binding:"required"` // 可能是用户名或邮箱
	PWD string `json:"pwd" binding:"required"`
}
type PWDLoginResponse struct {
	Token string `json:"token"`
}

func (UserApi) PWDLoginApi(c *gin.Context) {
	var req PWDLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithError(err, c)
		return
	}
	if !global.Config.Site.Login.UsernamePWDLogin {
		res.FailWithMsg("站点未启用用户名密码登录", c)
	}

	var m models.UserModel
	if err := global.DB.Take(&m, "(username = ? OR email = ?) AND password <> ''", req.Val, req.Val).Error; err != nil {
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	if !pwd.CompareHashAndPassword(m.Password, req.PWD) {
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	// 颁发token
	claims := jwts.Claims{
		UserID:   m.ID,
		Username: m.Username,
		Role:     m.Role,
	}
	token, err := jwts.GetToken(claims)
	if err != nil {
		logrus.Errorf("JWT生成Token失败 %s", err)
		res.FailWithMsg("登录失败", c)
		return
	}

	v := PWDLoginResponse{Token: token}
	res.OKWithData(v, c)
}
