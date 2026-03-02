package device

import "log/slog"

type Imei string

type Device struct {
	Imei   Imei
	Serial uint

	State *state

	Gps *gpsInfo
	Lbs *lbsInfo

	logger *slog.Logger
}

func NewDevice(imei Imei) *Device {
	return &Device{
		Imei:   imei,
		Serial: 0,
	}
}
