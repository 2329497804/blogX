package conf

type Config struct {
	Server Server `yaml:"server"`
	Log    Log    `yaml:"log"`
}
