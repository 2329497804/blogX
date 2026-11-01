package core

import (
	"fmt"
	"os"

	"github.com/2329497804/blogX-server/conf"
	"github.com/2329497804/blogX-server/flags"
	"gopkg.in/yaml.v3"
)

func ReadConfig() (c *conf.Config) {
	byteData, err := os.ReadFile(flags.FlagOptions.File)
	if err != nil {
		panic(err)
	}
	c = new(conf.Config)
	err = yaml.Unmarshal(byteData, c)
	if err != nil {
		panic(fmt.Sprintf("yaml Unmarshal err: %v", err))
	}
	fmt.Printf("读取配置文件 %s 成功\n", flags.FlagOptions.File)
	return
}
