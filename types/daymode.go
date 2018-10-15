package types

type DayMode string

const (
	DayModeDaysOfWeek  = "DaysOfWeek"
	DayModeDayInterval = "Interval"
)

var dayModeMap = map[int]DayMode{
	0: DayModeDaysOfWeek,
	1: DayModeDayInterval,
}

func GetDayMode(value int) DayMode {
	if val, ok := dayModeMap[value]; ok {
		return val
	}

	return "Unknown"
}
