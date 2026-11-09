package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func init() {
	// 创建翻译器
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")

	// 注册翻译器
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		panic("validator: unsupported validator engine")
	}
	_ = zh_translations.RegisterDefaultTranslations(v, trans)

	// 字段名也显示中文(用户体验更好)
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			label = field.Name
		}
		name := field.Tag.Get("json")
		return fmt.Sprintf("%s---%s", name, label)
	})
}

func ValidateError(err error) (data map[string]any, msg string) {
	var errs validator.ValidationErrors
	if ok := errors.As(err, &errs); !ok {
		msg = err.Error()
		return
	}
	data = make(map[string]any)
	var msgList []string
	for _, e := range errs {
		m := e.Translate(trans)
		_list := strings.Split(m, "---")
		data[_list[0]] = _list[1]
		msgList = append(msgList, _list[1])
	}
	msg = strings.Join(msgList, ";")
	return
}
