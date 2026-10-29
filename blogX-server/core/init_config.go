package core

import (
	"fmt"
	"os"

	"github.com/2329497804/blogX-server/flags"
	"gopkg.in/yaml.v3"
)

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}
type Config struct {
	Server Server `yaml:"server"`
}

func ReadConfig() {
	byteData, err := os.ReadFile(flags.FlagOptions.File)
	if err != nil {
		panic(err)
	}
	var conf Config
	err = yaml.Unmarshal(byteData, &conf)
	if err != nil {
		panic(fmt.Sprintf("yaml Unmarshal err: %v", err))
	}
	fmt.Printf("读取配置文件 %s 成功", flags.FlagOptions.File)
}
