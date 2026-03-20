package device

import (
	"go_tcp-tracker-simulator/internal/domain"
	"go_tcp-tracker-simulator/internal/protocol"
	"time"
)

type DeviceEvent struct {
	Kind domain.DeviceEventType

	IMEI      domain.IMEI
	SessionID domain.SessionID

	Packet protocol.Packet
	Error  error

	Time time.Time
}
