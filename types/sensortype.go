package types

type SensorType string

const (
	SensorTypeNone           = "None"
	SensorTypeTemperature    = "Temperature"
	SensorTypePH             = "PH"
	SensorTypeRedox          = "Redox"
	SensorTypeConductivityF  = "FreshwaterConductivity"
	SensorTypeConductivity   = "Conductivity"
	SensorTypeFree           = "Free"
	SensorTypeHumidity       = "Humidity"
	SensorTypeAirTemperature = "AirTemperature"
	SensorTypeOxygen         = "Oxygen"
	SensorTypeVoltage        = "Voltage"
)

var sensorTypeMap = map[int]string{
	0:  SensorTypeNone,
	1:  SensorTypeTemperature,
	2:  SensorTypePH,
	3:  SensorTypeRedox,
	4:  SensorTypeConductivityF,
	5:  SensorTypeConductivity,
	6:  SensorTypeFree,
	7:  SensorTypeHumidity,
	8:  SensorTypeAirTemperature,
	9:  SensorTypeOxygen,
	10: SensorTypeVoltage,
}

func GetSensorType(value int) string {
	if val, ok := sensorTypeMap[value]; ok {
		return val
	}

	return "Unknown"
}
