package enum

type LogLevelType int8

const (
	_ LogLevelType = iota
	InfoLogLevelType
	WarnLogLevelType
	ErrorLogLevelType
)
