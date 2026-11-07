package middleware

import (
	"net/http"

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
		Head:           make(http.Header), //防止空指针
	}
	c.Writer = res
	c.Next()
	//响应中间件
	log.SetResponseHeader(res.Head)
	log.SetResponseBody(res.Body)
	log.Save()
}

type ResponseWriter struct {
	gin.ResponseWriter
	Head http.Header
	Body []byte
}

func (r *ResponseWriter) Header() http.Header {
	return r.Head
}

func (r *ResponseWriter) Write(b []byte) (int, error) {
	r.Body = append(r.Body, b...)
	return r.ResponseWriter.Write(b)
}
