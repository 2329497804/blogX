package models

// BannerModel banner表
type BannerModel struct {
	Model
	Show  bool   `json:"show"`  // 是否展示
	Cover string `json:"cover"` // 图片链接
	Href  string `json:"href"`  // 跳转链接
}
