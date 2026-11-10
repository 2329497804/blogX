package middleware

import (
	"github.com/2329497804/blogX-server/common/res"
	"github.com/2329497804/blogX-server/models/enum"
	"github.com/2329497804/blogX-server/service/redis_service/redis_jwt"
	"github.com/2329497804/blogX-server/utils/jwts"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	claims, ok := authenticate(c)
	if !ok {
		return
	}
	c.Set("claims", claims)
	return
}

func AdminMiddleware(c *gin.Context) {
	var claims jwts.MyClaims

	if v, ok := c.Get("claims"); ok {
		// 已经经过AuthMiddleware 复用
		claims = v.(jwts.MyClaims)
	} else {
		var ok bool
		claims, ok = authenticate(c)
		if !ok {
			return
		}
		c.Set("claims", claims)
	}
	if claims.Role != enum.AdminRole {
		res.FailWithMsg("权限错误", c)
		c.Abort()
		return
	}
	return
}

func authenticate(c *gin.Context) (jwts.MyClaims, bool) {
	t, ok := redis_jwt.HasTokenBlackByGin(c)
	if ok {
		res.FailWithMsg(t.Msg(), c)
		c.Abort()
		return jwts.MyClaims{}, false
	}

	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return jwts.MyClaims{}, false
	}

	return claims, true
}
