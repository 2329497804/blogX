package main

import (
	"github.com/2329497804/blogX-server/core"
	"github.com/2329497804/blogX-server/flags"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/router"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConfig()
	core.InitLogrus()
	global.DB = core.InitDB()
	global.Redis = core.InitRedis()

	flags.Run()

	//启动Web程序
	router.Run()
}
