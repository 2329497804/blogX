package flags

import (
	"fmt"
	"os"

	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models"
	"github.com/2329497804/blogX-server/models/enum"
	"github.com/2329497804/blogX-server/utils/pwd"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

type FlagUser struct {
}

func (FlagUser) Create() {
	var role enum.RoleType
	fmt.Printf("选择角色: %d管理员, %d用户, %d访客\n", enum.AdminRole, enum.UserRole, enum.VisitorRole)
	_, err := fmt.Scan(&role)
	if err != nil {
		logrus.Errorf("输入错误 %s", err)
		return
	}
	if !(role == enum.AdminRole || role == enum.UserRole || role == enum.VisitorRole) {
		logrus.Errorf("输入角色不存在")
		return
	}

	var username string
	fmt.Println("请输入用户名: ")
	fmt.Scan(&username)
	// 查用户名是否存在
	var m models.UserModel
	err = global.DB.Take(&m, "username = ?", username).Error
	if err == nil {
		logrus.Errorf("用户名已存在")
		return
	}

	var pwd1, pwd2 []byte
	fmt.Println("请输入密码: ")
	pwd1, err = terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		logrus.Errorf("读取密码出错: %s", err)
		return
	}
	fmt.Println("请再次确认密码: ")
	pwd2, err = terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		logrus.Errorf("读取密码出错: %s", err)
		return
	}
	if string(pwd1) != string(pwd2) {
		logrus.Errorf("两次输入密码不一致")
		return
	}
	hashPWD, err := pwd.GenerateFromPWD(string(pwd1))

	// 创建用户
	err = global.DB.Create(&models.UserModel{
		Username:       username,
		Nickname:       username,
		RegisterSource: enum.TerminalRegisterType,
		Password:       hashPWD,
		Role:           role,
	}).Error
	if err != nil {
		logrus.Errorf("创建用户失败: %s", err)
		return
	}
	logrus.Infof("创建用户成功")
}
