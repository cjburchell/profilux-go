package types

type LogicMode string

const (
	LogicModeAnd            = "And"
	LogicModeOr             = "Or"
	LogicModeInvertAnd      = "InvertAnd"
	LogicModeInvertOr       = "InvertOr"
	LogicModeInverted       = "Inverted"
	LogicModeEqual          = "Equal"
	LogicModeNoEqual        = "NoEqual"
	LogicModePulse          = "Pulse"
	LogicModeDelayedOn      = "DelayedOn"
	LogicModeDelayedOff     = "DelayedOff"
	LogicModeFrequentPulses = "FrequentPulses"
	LogicModeSRFlipFlop     = "SRFlipFlop"
	LogicModeExclusiveOr    = "ExclusiveOr"
)

var logicModeMap = map[int]LogicMode{
	0:  LogicModeAnd,
	1:  LogicModeOr,
	2:  LogicModeInvertAnd,
	3:  LogicModeInvertOr,
	4:  LogicModeInverted,
	5:  LogicModeEqual,
	6:  LogicModeNoEqual,
	7:  LogicModePulse,
	8:  LogicModeDelayedOn,
	9:  LogicModeDelayedOff,
	10: LogicModeFrequentPulses,
	11: LogicModeSRFlipFlop,
	12: LogicModeExclusiveOr,
}

func GetLogicMode(value int) LogicMode {
	if val, ok := logicModeMap[value]; ok {
		return val
	}

	return "Unknown"
}
