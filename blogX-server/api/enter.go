package api

import "github.com/2329497804/blogX-server/api/site_api"

type Api struct {
	SiteApi site_api.SiteApi
}

var App = Api{}
