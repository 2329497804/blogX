package common

import (
	"fmt"

	"github.com/2329497804/blogX-server/global"
	"gorm.io/gorm"
)

type PageInfo struct {
	Limit int    `form:"limit"`
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Order string `form:"order"` //前端可覆盖
}

func (p PageInfo) GetPage() int {
	if p.Page > 20 || p.Page < 1 {
		return 1
	}
	return p.Page
}
func (p PageInfo) GetLimit() int {
	if p.Limit > 100 || p.Limit < 1 {
		return 10
	}
	return p.Limit
}
func (p PageInfo) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

type Options struct {
	PageInfo
	Likes        []string
	Preload      []string
	Where        *gorm.DB
	Debug        bool
	DefaultOrder string
}

func ListQuery[T any](model T, option Options) (list []T, count int, err error) {
	// 基础查询
	query := global.DB.Model(model).Where(model)

	//日志
	if option.Debug {
		query = query.Debug()
	}

	// 模糊匹配
	if len(option.Likes) > 0 && option.Key != "" {
		like := global.DB.Where("")
		for _, column := range option.Likes {
			like = like.Or(fmt.Sprintf("%s LIKE ?", column), "%"+option.Key+"%")
		}
		query = query.Where(like)
	}

	// 预加载
	if len(option.Preload) > 0 {
		for _, v := range option.Preload {
			query = query.Preload(v)
		}
	}

	// 定制化查询
	if option.Where != nil {
		query = query.Where(option.Where)
	}

	// 总数
	var _c int64
	query.Count(&_c)
	count = int(_c)

	// 分页
	l := option.GetLimit()
	o := option.GetOffset()
	query = query.Offset(o).Limit(l)

	//排序
	if option.Order != "" {
		query = query.Order(option.Order)
	} else {
		if option.DefaultOrder != "" {
			query = query.Order(option.DefaultOrder)
		}
	}

	err = query.Find(&list).Error
	return
}
