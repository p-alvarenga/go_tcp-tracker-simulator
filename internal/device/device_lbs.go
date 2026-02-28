package device

type lbsInfo struct {
	Mcc    uint16 // Mobile Country Code
	Mnc    uint8  // Mobile Network Code
	Lac    uint16 // Location Area Code
	CellId uint32
}
