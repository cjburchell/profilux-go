package types

type SensorMode string

const (
	SensorModeNormal  = "Normal"
	SensorModeAverage = "Average"
	SensorModeCopy    = "Copy"
)

var sensorModeMap = map[int]string{
	0: SensorModeNormal,
	1: SensorModeAverage,
	4: SensorModeCopy,
}

func GetSensorMode(value int) string {
	if val, ok := sensorModeMap[value]; ok {
		return val
	}

	return "Unknown"
}
