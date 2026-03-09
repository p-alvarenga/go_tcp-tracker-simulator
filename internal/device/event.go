package device

import (
	"go_tcp-tracker-simulator/internal/domain"
	"time"
)

type DeviceEvent struct {
	Kind domain.DeviceEventType

	IMEI      domain.IMEI
	SessionID domain.SessionID

	Login *LoginPayload

	Time time.Time
}

func NewDeviceEvent() *DeviceEvent {
	return nil
}
