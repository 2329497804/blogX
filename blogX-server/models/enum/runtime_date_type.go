package enum

type RuntimeDateType int8

// 按日期类型分割运行日志
const (
	_ RuntimeDateType = iota
	RuntimeDateHour
	RuntimeDateDay
	RuntimeDateWeek
	RuntimeDateMonth
)

func (r RuntimeDateType) GetSQLTime() string {
	switch r {
	case RuntimeDateHour:
		return "interval 1 HOUR"
	case RuntimeDateDay:
		return "interval 1 DAY"
	case RuntimeDateWeek:
		return "interval 1 WEEK"
	case RuntimeDateMonth:
		return "interval 1 MONTH"
	}
	return "interval 1 DAY"
}
