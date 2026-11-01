package core

import (
	"time"

	"github.com/2329497804/blogX-server/global"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	dc := global.Config.DB

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dc.DSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 不生成外键约束
	})
	if err != nil {
		logrus.Fatalf("init db err: %s", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logrus.Fatalf("init db err: %s", err)
	}
	sqlDB.SetMaxIdleConns(10)           // SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	sqlDB.SetMaxOpenConns(100)          // SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(time.Hour) // SetConnMaxLifetime 设置了可以重新使用连接的最大时间。
	logrus.Infof("init db success")
	return db
}
