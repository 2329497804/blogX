package utils

import (
	"errors"
	"fmt"
	"strings"

	e "github.com/pkg/errors"
)

type stackTracer interface {
	error
	StackTrace() e.StackTrace
}

func ErrorStack(err error) string {
	if err == nil {
		return ""
	}

	var st stackTracer
	if !errors.As(err, &st) {
		return err.Error()
	}

	var b strings.Builder
	b.WriteString(err.Error())

	for _, frame := range st.StackTrace() {
		fn := fmt.Sprintf("%+v", frame)
		if isNoiseFrame(fn) {
			continue
		}

		b.WriteString("\n")
		b.WriteString(fn)
	}

	return b.String()
}

// 不属于业务模块的栈帧前缀
var noisePrefixes = []string{
	"github.com/gin-gonic/gin.", //小数点匹配包下的函数
	"github.com/gin-contrib/",
	"net/http.",
	"runtime.",
}

func isNoiseFrame(fn string) bool {
	for _, prefix := range noisePrefixes {
		if strings.HasPrefix(fn, prefix) {
			return true
		}
	}
	return false
}
