package types

type LevelSensorOperationMode string

const (
	LevelSensorNotEnabled               = "NotEnabled"
	LevelSensorAutoTopOff               = "AutoTopOff"
	LevelSensorMinMaxControl            = "MinMaxControl"
	LevelSensorWaterChange              = "WaterChange"
	LevelSensorLeakageDetection         = "LeakageDetection"
	LevelSensorWaterChangeAndAutoTopOff = "WaterChangeAndAutoTopOff"
	LevelSensorAutoTopOffWith2Sensors   = "AutoTopOffWith2Sensors"
	LevelSensorReturnPump               = "ReturnPump"
)

var levelSensorOperationModeMap = map[int]string{
	0: LevelSensorNotEnabled,
	1: LevelSensorAutoTopOff,
	2: LevelSensorMinMaxControl,
	3: LevelSensorWaterChange,
	4: LevelSensorLeakageDetection,
	5: LevelSensorWaterChangeAndAutoTopOff,
	6: LevelSensorAutoTopOffWith2Sensors,
	7: LevelSensorReturnPump,
}

func GetLevelSensorOperationMode(value int) string {
	if val, ok := levelSensorOperationModeMap[value]; ok {
		return val
	}

	return "Unknown"
}
