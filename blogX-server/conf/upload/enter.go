package upload

type Image struct {
	Size      int64    `yaml:"size"`
	Dir       string   `yaml:"dir"`
	WhiteList []string `yaml:"white_list"`
}
