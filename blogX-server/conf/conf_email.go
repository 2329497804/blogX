package conf

type Email struct {
	Domain       string `yaml:"domain" json:"domain"`
	Port         string `yaml:"port" json:"port"`
	SendEmail    string `yaml:"send_email" json:"sendEmail"`
	AuthCode     string `yaml:"auth_code" json:"authCode"` // 授权码
	SendNickName string `yaml:"send_nickname" json:"sendName"`
	SSL          bool   `yaml:"ssl" json:"ssl"`
	TLS          bool   `yaml:"tls" json:"tls"`
}
