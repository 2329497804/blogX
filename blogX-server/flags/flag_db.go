package flags

import (
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models"
	"github.com/sirupsen/logrus"
)

func FlagDB() {
	err := global.DB.AutoMigrate(
		&models.UserModel{},
		&models.UserConfModel{},
	)
	if err != nil {
		logrus.Errorf("AutoMigrate err: %v", err)
		return
	}
	logrus.Info("AutoMigrate success")
}
