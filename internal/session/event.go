package session

import (
	"go_tcp-tracker-simulator/internal/domain"
	"time"
)

type SessionEvent struct {
	Kind      domain.SessionEventType
	IMEI      domain.IMEI
	SessionID domain.SessionID

	Error error
	Time  time.Time
}
