package redis_jwt

import (
	"fmt"
	"time"

	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/utils/jwts"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type BlackType int8

const (
	_               BlackType = iota
	UserBlackType             // 用户注销登录
	AdminBlackType            // 管理员强制下线
	DeviceBlackType           // 其他设备挤下线
)

func (b BlackType) String() string {
	return fmt.Sprintf("%d", b)
}
func (b BlackType) Msg() string {
	switch b {
	case UserBlackType:
		return "已注销"
	case AdminBlackType:
		return "禁止登录"
	case DeviceBlackType:
		return "设备下线"
	}
	return "禁止登录"
}
func ParseBlackType(s string) BlackType {
	switch s {
	case "1":
		return UserBlackType
	case "2":
		return AdminBlackType
	case "3":
		return DeviceBlackType
	}
	return UserBlackType
}

func TokenBlack(token string, value BlackType) {
	key := fmt.Sprintf("token_black_%s", token)

	claims, err := jwts.ParseToken(token)
	if err != nil || claims == nil {
		logrus.Errorf("token解析失败: %s", err)
		return
	}
	expire := time.Duration(claims.ExpiresAt-time.Now().Unix()) * time.Hour
	_, err = global.Redis.Set(key, value.String(), expire).Result()
	if err != nil {
		logrus.Errorf("redis添加黑名单失败 %s", err)
		return
	}
}

func HasTokenBlack(token string) (b BlackType, ok bool) {
	key := fmt.Sprintf("token_black_%s", token)
	res, err := global.Redis.Get(key).Result()
	if err != nil {
		return
	}
	return ParseBlackType(res), true
}

func HasTokenBlackByGin(c *gin.Context) (BlackType, bool) {
	token := jwts.GetTokenByGin(c)
	return HasTokenBlack(token)
}
