package global

import (
	"github.com/2329497804/blogX-server/conf"
	"gorm.io/gorm"
)

var (
	Config *conf.Config
	DB     *gorm.DB
)
