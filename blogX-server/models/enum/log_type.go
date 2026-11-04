package enum

type LogType int8

const (
	_ LogType = iota
	LoginLogType
	ActionLogType
	RuntimeLogType
)
