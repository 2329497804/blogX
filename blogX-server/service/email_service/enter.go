package email_service

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/2329497804/blogX-server/global"
	"github.com/jordan-wright/email"
)

// SendRegisterCode 注册账号邮件验证码
func SendRegisterCode(to, code string) error {
	c := global.Config.Email
	sub := fmt.Sprintf("【%s】账号注册", c.SendNickName)
	text := fmt.Sprintf("你正在进行账号注册, 验证码: %s (10分钟内有效)", code)
	return SendEmail(to, sub, text)
}

// SendResetPWDCode 重置密码邮件验证码
func SendResetPWDCode(to, code string) error {
	c := global.Config.Email
	sub := fmt.Sprintf("【%s】密码重置", c.SendNickName)
	text := fmt.Sprintf("你正在进行密码重置, 验证码: %s (10分钟内有效)", code)
	return SendEmail(to, sub, text)
}

func SendEmail(to, sub, text string) error {
	c := global.Config.Email
	e := &email.Email{
		To:      []string{to},
		From:    fmt.Sprintf("%s <%s>", c.SendNickName, c.SendEmail),
		Subject: sub,
		Text:    []byte(text),
	}
	addr := fmt.Sprintf("%s:%s", c.Domain, c.Port)
	err := e.Send(addr, smtp.PlainAuth("", c.SendEmail, c.AuthCode, c.Domain))
	if err != nil && !strings.Contains(err.Error(), "short response:") {
		return err
	}
	return nil
}
