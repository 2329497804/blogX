package user_api

import (
	"math"
	"time"

	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/models"
	"github.com/2329497804/blogX-server/models/enum"
	"github.com/2329497804/blogX-server/utils/jwts"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserDetailResponse struct {
	ID                 uint              `json:"id"`
	Username           string            `json:"username"`
	Nickname           string            `json:"nickname,"`
	Avatar             string            `json:"avatar"`
	Abstract           string            `json:"abstract"`
	RegisterSource     enum.RegisterType `json:"registerSource"`
	LikeTags           []string          `json:"likeTags"`
	UpdateUsernameDate *time.Time        `json:"updateUsernameDate"`
	OpenCollect        bool              `json:"openCollect"`
	OpenFollow         bool              `json:"openFollow"`
	OpenFans           bool              `json:"openFans"`
	HomeStyleID        uint              `json:"homeStyleId"`
	CodeAge            int               `json:"codeAge"` // 码龄
	CreatedAt          time.Time         `json:"createdAt"`
}

func (UserApi) UserDetailView(c *gin.Context) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var m models.UserModel
	err = global.DB.Preload("UserConfModel").Take(&m, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	if m.UserConfModel == nil {
		logrus.Errorf("UserConfModel is nil, userID: %d", claims.UserID)
		res.FailWithMsg("用户配置为空", c)
		return
	}

	sub := time.Now().Sub(m.CreatedAt)
	codeAge := int(math.Ceil(sub.Hours() / 24 / 365))
	v := UserDetailResponse{
		ID:                 m.ID,
		Username:           m.Username,
		Nickname:           m.Nickname,
		Avatar:             m.Avatar,
		Abstract:           m.Abstract,
		RegisterSource:     m.RegisterSource,
		LikeTags:           m.UserConfModel.LikeTags,
		UpdateUsernameDate: m.UserConfModel.UpdateUsernameDate,
		OpenCollect:        m.UserConfModel.OpenCollect,
		OpenFollow:         m.UserConfModel.OpenFollow,
		OpenFans:           m.UserConfModel.OpenFans,
		HomeStyleID:        m.UserConfModel.HomeStyleID,
		CreatedAt:          m.CreatedAt,
		CodeAge:            codeAge,
	}

	res.OKWithData(v, c)
}
