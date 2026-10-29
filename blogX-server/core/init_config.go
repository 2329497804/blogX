package core

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type Config struct {
	Server Server `yaml:"server"`
}

var confPath = "settings.yaml"

func ReadConfig() {
	byteData, err := os.ReadFile(confPath)
	if err != nil {
		panic(err)
	}
	var conf Config
	err = yaml.Unmarshal(byteData, &conf)
	if err != nil {
		panic(fmt.Sprintf("yaml Unmarshal err: %v", err))
	}
	fmt.Println(conf)
}
