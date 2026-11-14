package conf

type QiNiu struct {
	Enable    bool   `yaml:"enable" json:"enable"`
	AccessKey string `yaml:"access_key" json:"accessKey"`
	SecretKey string `yaml:"secret_key" json:"secretKey"`
	Bucket    string `yaml:"bucket" json:"bucket"`
	URI       string `yaml:"uri" json:"uri"`
	Region    string `yaml:"region" json:"region"`
	Prefix    string `yaml:"prefix" json:"prefix"`
	Size      int    `yaml:"size" json:"size"` // 大小限制 单位mb
}
