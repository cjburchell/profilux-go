package types

type WaterMode string

const (
	WaterModeReady    = "Ready"
	WaterModeDraining = "Draining"
	WaterModeFilling  = "Filling"
)

var waterModeMap = map[int]WaterMode{
	0: WaterModeReady,
	1: WaterModeDraining,
	2: WaterModeFilling,
}

func GetWaterMode(value int) WaterMode {
	if val, ok := waterModeMap[value]; ok {
		return val
	}

	return "Unknown"
}
