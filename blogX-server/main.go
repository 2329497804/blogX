package main

import (
	"time"

	"github.com/2329497804/blogX-server/core"
	"github.com/2329497804/blogX-server/flags"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/router"
	"github.com/mojocn/base64Captcha"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConfig()
	core.InitLogrus()
	global.DB = core.InitDB()
	global.Redis = core.InitRedis()
	global.CaptchaStore = base64Captcha.NewMemoryStore(10240, 5*time.Minute)

	flags.Run()

	//启动Web程序
	router.Run()
}
