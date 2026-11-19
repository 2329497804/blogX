package enum

type RegisterType int8

const(
	_ RegisterType = iota
	EmailRegisterType
	QQRegisterType
	TerminalRegisterType
)

func (t RegisterType) String() string {
	switch t {
	case EmailRegisterType:
		return "Email"
	case QQRegisterType:
		return "QQ"
	case TerminalRegisterType:
		return "Terminal"
	}
	return "Unknown"
}