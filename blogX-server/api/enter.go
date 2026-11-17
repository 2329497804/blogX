package api

import (
	"github.com/2329497804/blogX-server/api/image_api"
	"github.com/2329497804/blogX-server/api/log_api"
	"github.com/2329497804/blogX-server/api/site_api"
)

type Api struct {
	SiteApi  site_api.SiteApi
	LogApi   log_api.LogApi
	ImageApi image_api.ImageApi
}

var App = Api{}
