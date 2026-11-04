package middleware

import (
	"github.com/2329497804/blogX-server/service/log_service"
	"github.com/gin-gonic/gin"
)

func LogMiddleWare(c *gin.Context) {
	//请求中间件
	log := log_service.NewActionLogByGin(c)
	c.Set("log", log)

	log.SetRequestBody(c)
	res := &ResponseWriter{
		ResponseWriter: c.Writer,
	}
	c.Writer = res
	c.Next()
	//响应中间件
	log.SetResponseBody(res.Body)
	log.Save()
}

type ResponseWriter struct {
	gin.ResponseWriter
	Body []byte
}

func (r *ResponseWriter) Write(b []byte) (int, error) {
	r.Body = append(r.Body, b...)
	return r.ResponseWriter.Write(b)
}
