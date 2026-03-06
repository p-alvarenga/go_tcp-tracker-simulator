package device

type lbs struct {
	MCC    uint16 // Mobile Country Code
	MNC    uint8  // Mobile Network Code
	LAC    uint16 // Location Area Code
	CellID uint32
}
