package enum

type LogLevelType int8

const (
	_ LogLevelType = iota
	InfoLogLevelType
	WarnLogLevelType
	ErrorLogLevelType
)

func (level LogLevelType) String() string {
	switch level {
	case InfoLogLevelType:
		return "info"
	case WarnLogLevelType:
		return "warn"
	case ErrorLogLevelType:
		return "error"
	}
	return ""
}
