package api

import (
	"github.com/2329497804/blogX-server/api/banner_api"
	"github.com/2329497804/blogX-server/api/captcha_api"
	"github.com/2329497804/blogX-server/api/image_api"
	"github.com/2329497804/blogX-server/api/log_api"
	"github.com/2329497804/blogX-server/api/site_api"
	"github.com/2329497804/blogX-server/api/user_api"
)

type Api struct {
	site_api.SiteApi
	log_api.LogApi
	image_api.ImageApi
	banner_api.BannerApi
	captcha_api.CaptchaApi
	user_api.UserApi
}

var App = Api{}
