package res

import (
	"github.com/2329497804/blogX-server/utils"
	"github.com/gin-gonic/gin"
)

type Code int

const (
	SuccessCode   Code = iota
	FailValidCode Code = 1000 + iota
	FailServiceCode
)

func (c Code) String() string {
	switch c {
	case SuccessCode:
		return "成功"
	case FailValidCode:
		return "校验失败"
	case FailServiceCode:
		return "服务异常"
	}
	return "未知状态"
}

type Response struct {
	Code Code   `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

var empty = map[string]any{}

func (r Response) Json(c *gin.Context) {
	c.JSON(200, r)
}

func OK(msg string, data any, c *gin.Context) {
	Response{SuccessCode, msg, data}.Json(c)
}
func OKWithMsg(msg string, c *gin.Context) {
	OK(msg, empty, c)
}
func OKWithData(data any, c *gin.Context) {
	OK(SuccessCode.String(), data, c)
}
func OKWithList(list any, count int, c *gin.Context) {
	Response{SuccessCode, SuccessCode.String(),
		map[string]any{
			"list":  list,
			"count": count,
		}}.Json(c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Response{FailValidCode, msg, empty}.Json(c)
}
func FailWithData(data any, msg string, c *gin.Context) {
	Response{FailServiceCode, msg, data}.Json(c)
}
func FailWithCode(code Code, c *gin.Context) {
	Response{code, code.String(), empty}.Json(c)
}
func FailWithError(err error, c *gin.Context) {
	data, msg := utils.ValidateError(err)
	Response{FailValidCode, msg, data}.Json(c)
}
