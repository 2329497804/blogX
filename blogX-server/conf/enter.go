package conf

type Config struct {
	Server Server `yaml:"server"`
	Log    Log    `yaml:"log"`
	DB     DB     `yaml:"db"`  // 读库
	DB1    DB     `yaml:"db1"` // 写库
	Jwt    Jwt    `yaml:"jwt"`
	Redis  Redis  `yaml:"redis"`
	Site   Site   `yaml:"site"`
	Email  Email  `yaml:"email"`
	QQ     QQ     `yaml:"qq"`
	QiNiu  QiNiu  `yaml:"qiniu"`
	AI     AI     `yaml:"ai"`
}
