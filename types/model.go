package types

import "fmt"

type Model string

const (
	ProductIdUnknown  = "Unknown"
	ProfiLuxII        = "ProfiLux II"
	ProfiLuxPlusII    = "ProfiLux Plus II"
	ProfiLuxPlusIIEx  = "ProfiLux Plus II Ex"
	ProfiLuxIITerra   = "ProfiLux II Terra"
	ProfiLuxIIEx      = "ProfiLux II Ex"
	ProfiLuxLightII   = "ProfiLux II Light"
	ProfiLuxIIOutdoor = "ProfiLux II Outdoor"
	ProfiLuxIII       = "ProfiLux III"
	ProfiLuxIIIEx     = "ProfiLux III Ex"
)

func GetModel(id int) string {
	if val, ok := productIdMap[id]; ok {
		return val
	}

	return fmt.Sprintf("Unknown (%d???)", id)
}

var productIdMap = map[int]string{
	0:  ProductIdUnknown,
	2:  ProfiLuxII,
	3:  ProfiLuxPlusII,
	4:  ProfiLuxPlusIIEx,
	5:  ProfiLuxIITerra,
	6:  ProfiLuxIIEx,
	7:  ProfiLuxLightII,
	8:  ProfiLuxIIOutdoor,
	11: ProfiLuxIII,
	12: ProfiLuxIIIEx,
}
