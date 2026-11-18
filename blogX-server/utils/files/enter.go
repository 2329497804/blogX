package files

import (
	"errors"
	"slices"
	"strings"

	"github.com/2329497804/blogX-server/global"
)

func ImageSuffixJudge(filename string) (suffix string, err error) {
	l := strings.Split(filename, ".")
	if len(l) < 2 {
		err = errors.New("错误的文件名")
		return
	}
	suffix = l[len(l)-1]
	whiteList := global.Config.Upload.Image.WhiteList
	ok := slices.Contains(whiteList, suffix)
	if !ok {
		err = errors.New("文件类型非法")
		return
	}
	return
}
