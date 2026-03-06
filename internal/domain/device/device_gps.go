package device

type gps struct {
	QuantityOfGpsInfo uint8
	RealTime          bool
	Trusful           bool
}

func newGPS(quantityOfGpsInfo uint8, realtime, trustful bool) *gps {
	return &gps{
		QuantityOfGpsInfo: quantityOfGpsInfo,
		RealTime:          realtime,
		Trusful:           trustful,
	}
}
