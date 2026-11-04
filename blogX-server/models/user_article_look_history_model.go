package models

// UserArticleLookHistoryModel 用户访问文章历史表
type UserArticleLookHistoryModel struct {
	Model
	UserID       uint         `json:"userID"`
	UserModel    UserModel    `gorm:"foreignKey:UserID" json:"-"`
	ArticleID    uint         `json:"articleID"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID" json:"-"`
}
