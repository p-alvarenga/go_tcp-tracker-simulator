package device

type state struct {
	Speed    uint
	Course   uint
	Ignition bool

	Lat float64
	Lng float64
}

func generateState() {
	// ...
}
