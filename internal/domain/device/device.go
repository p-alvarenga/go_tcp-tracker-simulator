package device

import "log/slog"

type IMEI string

type Device struct {
	IMEI   IMEI
	Serial int

	State *DeviceState

	GPS *gps
	LBS *lbs

	logger *slog.Logger
}

func NewDevice(imei IMEI, serial int) *Device {
	return &Device{
		IMEI:   imei,
		Serial: 0,
	}
}
