package device

type DeviceState struct {
	Speed    uint
	Course   uint
	Ignition bool

	Lat float64
	Lng float64
}
