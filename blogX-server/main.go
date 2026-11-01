package main

import (
	"github.com/2329497804/blogX-server/core"
	"github.com/2329497804/blogX-server/flags"
	"github.com/2329497804/blogX-server/global"
	"github.com/sirupsen/logrus"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConfig()
	core.InitLogrus()

	logrus.Warnf("xxx")
	logrus.Debugf("YYY")
	logrus.Errorf("ZZZ")
	logrus.Infof("1123")
}
