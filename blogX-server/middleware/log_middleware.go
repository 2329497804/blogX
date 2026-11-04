package middleware

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogMiddleWare(c *gin.Context) {
	//请求中间件
	bytedata, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Error(err.Error())
	}
	fmt.Printf("body: %s\n", string(bytedata))
	c.Request.Body = io.NopCloser(bytes.NewReader(bytedata))

	c.Next()
	//响应中间件
}
