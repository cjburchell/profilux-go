package types

type DigitalInputFunction string

const (
	DigitalInputFunctionNotUsed      = "NotUsed"
	DigitalInputFunctionLevelSensor  = "LevelSensor"
	DigitalInputFunctionWaterChange  = "WaterChange"
	DigitalInputFunctionMaintenance  = "Maintenance"
	DigitalInputFunctionFeedingPause = "FeedingPause"
	DigitalInputFunctionThunderstorm = "Thunderstorm"
)

func GetDigitalInputFunction(value int) string {

	function := (value >> 4) & 0xF
	def := value & 0xF

	switch function {
	case 1:
		return DigitalInputFunctionLevelSensor
	case 2:
		switch def {
		case 0:
			return DigitalInputFunctionWaterChange
		case 1:
			return DigitalInputFunctionMaintenance
		case 2:
			return DigitalInputFunctionFeedingPause
		case 3:
			return DigitalInputFunctionThunderstorm
		}
	}

	return DigitalInputFunctionNotUsed
}
