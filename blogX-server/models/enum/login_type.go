package enum

type LoginType int8

const (
	_ LoginType = iota
	UserPwdLoginType
	QQLoginType
	EmailLoginType
)
