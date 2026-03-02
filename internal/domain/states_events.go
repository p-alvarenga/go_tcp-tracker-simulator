package domain

import (
	"time"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain/device"
)

type SimulatedDeviceState int

const (
	StateNew SimulatedDeviceState = iota
	StateLoggedIn
)

type SimulatorEventType int

const (
	EventCreated SimulatorEventType = iota
	EventStarted
	EventStopped

	EventExpired
	EventConnected
	EventDisconnected
	EventReconnected
	EventLoginSucceeded
	EventLoginFailed

	EventLocationSucceeded
	EventLocationFailed

	EventProtocolViolation
	EventUnexpectedResponse
	EventInvalidStateTransition

	EventUnknown
)

type SimulatorEvent struct {
	Id   device.Imei
	Type SimulatorEventType
	Time time.Time
}
