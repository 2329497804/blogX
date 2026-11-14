package site

type SiteInfo struct {
	Title string `yaml:"title" json:"title"` // 网站标题
	Logo  string `yaml:"logo" json:"logo"`
	Beian string `yaml:"beian" json:"beian"`
	Mode  int8   `yaml:"mode" json:"mode"` // 1: 社区模式 2: 博客模式
}
type Project struct {
	Title   string `yaml:"title" json:"title"` // 网站title(与网站标题不同)
	Icon    string `yaml:"icon" json:"icon"`
	WebPath string `yaml:"web_path" json:"webPath"`
}
type Seo struct {
	Keywords    string `yaml:"keywords" json:"keywords"`
	Description string `yaml:"description" json:"description"`
}
type About struct {
	SiteDate string `yaml:"site_date" json:"siteDate"` // 年月日
	QQ       string `yaml:"qq" json:"QQ"`
	Wechat   string `yaml:"wechat" json:"wechat"`
	Gitee    string `yaml:"gitee" json:"gitee"`
	Bilibili string `yaml:"bilibili" json:"bilibili"`
	Github   string `yaml:"github" json:"github"`
}
type Login struct {
	QQLogin          bool `yaml:"qq_login" json:"QQLogin"`
	UsernamePWDLogin bool `yaml:"username_pwd_login" json:"usernamePWDLogin"`
	EmailLogin       bool `yaml:"email_login" json:"emailLogin"`
	Captcha          bool `yaml:"captcha" json:"captcha"`
}
type ComponentInfo struct {
	Title  string `yaml:"title" json:"title"`
	Enable bool   `yaml:"enable" json:"enable"`
}
type IndexRight struct {
	List []ComponentInfo `yaml:"list" json:"list"`
}
type Article struct {
	NoExamine bool `yaml:"no_examine" json:"noExamine"` // 免审核
}
