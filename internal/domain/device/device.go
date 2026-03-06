package device

import "log/slog"

type IMEI string

type Device struct {
	IMEI   IMEI
	Serial int

	State *state

	Gps *gpsInfo
	Lbs *lbsInfo

	logger *slog.Logger
}

func NewDevice(imei IMEI, serial int) *Device {
	return &Device{
		IMEI:   imei,
		Serial: 0,
	}
}
