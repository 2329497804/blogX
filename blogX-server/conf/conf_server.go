package conf

import "fmt"

type Server struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Env     string `yaml:"env"`
	GinMode string `yaml:"gin_mode"`
}

func (s *Server) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
