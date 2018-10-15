package types

type TimerMode string

const (
	TimerModeNormal       = "Normal"
	TimerModeShort        = "Short"
	TimerModeAutoDosing   = "AutoDosing"
	TimerModeManualDosing = "ManualDosing"
	TimerModeStartEvent   = "StartEvent"
	TimerModeCyclic       = "Cyclic"
)

var timerModeMap = map[int]TimerMode{
	0: TimerModeNormal,
	1: TimerModeShort,
	2: TimerModeAutoDosing,
	3: TimerModeManualDosing,
	4: TimerModeStartEvent,
	5: TimerModeCyclic,
}

func GetTimerMode(value int) TimerMode {
	if val, ok := timerModeMap[value]; ok {
		return val
	}

	return "Unknown"
}
