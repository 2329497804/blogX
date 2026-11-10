package core

import (
	"fmt"
	"strings"

	ipUtils "github.com/2329497804/blogX-server/utils/ip"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/sirupsen/logrus"
)

var searcher *xdb.Searcher

func InitIPDB() {
	var version = xdb.IPv4
	var dbPath = "init/ip2region.xdb"
	_searcher, err := xdb.NewWithFileOnly(version, dbPath)
	if err != nil {
		logrus.Fatalf("failed to create searcher: %s\n", err.Error())
	}
	searcher = _searcher
}

func GetIPAddr(ip string) string {
	if ipUtils.HasLocalIPAddr(ip) {
		return "内网"
	}
	region, err := searcher.Search(ip)
	if err != nil {
		logrus.Warnf("错误的ip地址 %s", err)
		return "异常地址"
	}
	_addrList := strings.Split(region, "|")
	if len(_addrList) != 5 {
		logrus.Warnf("异常的ip地址 %s", ip)
		return "未知地址"
	}
	// _addrList 五个部分
	// 国家  0  省份   市   运营商
	country := _addrList[0]
	province := _addrList[2]
	city := _addrList[3]
	if province != "0" && city != "0" {
		return fmt.Sprintf("%s·%s", province, city)
	}
	if country != "0" && province != "0" {
		return fmt.Sprintf("%s·%s", country, province)
	}
	if country != "0" {
		return country
	}
	return region
}
