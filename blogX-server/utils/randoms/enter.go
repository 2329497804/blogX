package randoms

import "github.com/mojocn/base64Captcha"

// CodeNum4 生成 4位验证码
func CodeNum4() string {
	return Code(4, "0123456789")
}

// CodeNum5 生成 5位验证码
func CodeNum5() string {
	return Code(5, "0123456789")
}

// CodeNum6 生成 6位验证码
func CodeNum6() string {
	return Code(6, "0123456789")
}

func Code(size int, chars string) string {
	return base64Captcha.RandText(size, chars)
}

// ID 生成随机id
func ID() string {
	return base64Captcha.RandomId()
}
