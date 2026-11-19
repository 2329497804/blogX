package user_api

import (
	"fmt"

	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models"
	"github.com/2329497804/blogX-server/models/enum"
	"github.com/2329497804/blogX-server/service/email_service"
	"github.com/2329497804/blogX-server/utils/emails"
	"github.com/2329497804/blogX-server/utils/jwts"
	"github.com/2329497804/blogX-server/utils/pwd"
	"github.com/2329497804/blogX-server/utils/randoms"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	_ int8 = iota
	typeRegister
	typeResetPWD
)

type SendEmailRequest struct {
	Type  int8   `json:"type" binding:"required"` // 1 注册 2 重置密码
	Email string `json:"email" binding:"required"`
}
type SendEmailResponse struct {
	ID string `json:"id"`
}

func (UserApi) SendEmailView(c *gin.Context) {
	var req SendEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithError(err, c)
		return
	}
	code := randoms.CodeNum5()
	id := randoms.ID()
	var err error
	switch req.Type {
	case typeRegister:
		// 查询邮箱是否注册
		var m models.UserModel
		err = global.DB.Take(&m, "email = ?", req.Email).Error
		if err == nil { // 邮箱已存在
			res.FailWithMsg("该邮箱已绑定", c)
			return
		}
		err = email_service.SendRegisterCode(req.Email, code)
	case typeResetPWD:
		err = email_service.SendResetPWDCode(req.Email, code)
	}
	if err != nil {
		logrus.Errorf("邮件发送失败 %s", err)
		res.FailWithMsg("邮件发送失败", c)
		return
	}
	// 暂存邮箱及其验证码
	global.EmailStore.Store(id, emails.EmailStoreInfo{
		Email: req.Email,
		Code:  code,
	})

	v := SendEmailResponse{ID: id}
	res.OKWithData(v, c)
}

type RegisterEmailRequest struct {
	PWD string `json:"pwd" binding:"required"`
}
type RegisterEmailResponse struct {
	Token string `json:"token"`
}

func (UserApi) RegisterUserEmailView(c *gin.Context) {
	var req RegisterEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithError(err, c)
		return
	}

	// 创建用户
	uname := fmt.Sprintf("E_%s", randoms.CodeNum6())
	nname := fmt.Sprintf("邮箱用户%s", randoms.CodeNum4())
	hashedPWD, err := pwd.GenerateFromPWD(req.PWD)
	if err != nil {
		logrus.Errorf("生成密码哈希失败 %s", err)
		res.FailWithMsg("邮箱注册失败", c)
		return
	}
	_email, _ := c.Get("email")
	email := _email.(string)
	u := models.UserModel{
		Username:       uname,
		Nickname:       nname,
		RegisterSource: enum.EmailRegisterType,
		Password:       hashedPWD,
		Email:          email,
		Role:           enum.UserRole,
	}
	err = global.DB.Create(&u).Error
	if err != nil {
		logrus.Errorf("创建用户失败 %s", err)
		res.FailWithMsg("邮箱注册失败", c)
		return
	}

	// 颁发token
	claims := jwts.Claims{
		UserID:   u.ID,
		Username: u.Username,
		Role:     u.Role,
	}
	token, err := jwts.GetToken(claims)
	if err != nil {
		logrus.Errorf("JWT生成Token失败 %s", err)
		res.FailWithMsg("邮箱登录失败", c)
		return
	}

	v := RegisterEmailResponse{Token: token}
	res.OKWithData(v, c)
}
