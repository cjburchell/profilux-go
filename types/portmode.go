package types

type PortMode struct {
	DeviceMode DeviceMode
	Port       int
	BlackOut   int
	Invert     bool
	Id         string
	IsProbe    bool
}
