package device

type gpsInfo struct {
	QuantityOfGpsInfo uint8
	RealTime          bool
	Trusful           bool
}

func newGPSInfo(quantityOfGpsInfo uint8, realtime, trustful bool) *gpsInfo {
	return &gpsInfo{
		QuantityOfGpsInfo: quantityOfGpsInfo,
		RealTime:          realtime,
		Trusful:           trustful,
	}
}
