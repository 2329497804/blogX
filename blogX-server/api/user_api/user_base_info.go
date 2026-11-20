package user_api

import (
	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models"
	"github.com/gin-gonic/gin"
)

type UserBaseInfoResponse struct {
	UserID       uint   `json:"userID"`
	CodeAge      int    `json:"codeAge"`
	Avatar       string `json:"avatar"`
	Nickname     string `json:"nickname"`
	LookCount    int    `json:"lookCount"`
	ArticleCount int    `json:"articleCount"`
	FansCount    int    `json:"fansCount"`
	FollowCount  int    `json:"followCount"`
	Place        string `json:"place"`
}

func (UserApi) UserBaseInfoView(c *gin.Context) {
	var req models.IDRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		res.FailWithError(err, c)
		return
	}
	var u models.UserModel
	err := global.DB.Take(&u, req.ID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	data := UserBaseInfoResponse{
		UserID:       u.ID,
		CodeAge:      u.CodeAge(),
		Avatar:       u.Avatar,
		Nickname:     u.Nickname,
		LookCount:    0, // TODO: 需先做完文章
		ArticleCount: 0,
		FansCount:    0, // TODO: 需先做完好友关系
		FollowCount:  0,
		Place:        u.Addr,
	}

	res.OKWithData(data, c)
}
