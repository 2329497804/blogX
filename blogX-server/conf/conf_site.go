package conf

import "github.com/2329497804/blogX-server/conf/site"

type Site struct {
	site.SiteInfo   `yaml:"site_info" json:"siteInfo"`
	site.Project    `yaml:"project" json:"project"`
	site.Seo        `yaml:"seo" json:"seo"`
	site.About      `yaml:"about" json:"about"`
	site.Login      `yaml:"login" json:"login"`
	site.IndexRight `yaml:"index_right" json:"indexRight"`
	site.Article    `yaml:"article" json:"article"`
}
