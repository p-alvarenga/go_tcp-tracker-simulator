package domain

import (
	"fmt"
	"time"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain/device"
)

type SimulatedDeviceState int

const (
	StateNew SimulatedDeviceState = iota
	StateLoggedIn
)

var simDeviceStateNames = map[SimulatedDeviceState]string{
	StateNew:      "STATE_NEW",
	StateLoggedIn: "STATE_LOGGED_IN",
}

func (s SimulatedDeviceState) String() string {
	if !s.IsValid() {
		return fmt.Sprintf("SIMULATED_DEVICE_STATE(%d)", int(s))
	}

	return simDeviceStateNames[s]
}

func (s SimulatedDeviceState) IsValid() bool {
	return s >= 0 && int(s) < len(simDeviceStateNames)
}

// SIMULATOR EVENT TYPE

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

	EventInvalid
)

var simulatorEventTypeNames = map[SimulatorEventType]string{
	EventCreated:                "EVENT_CREATED",
	EventStarted:                "EVENT_STARTED",
	EventStopped:                "EVENT_STOPPED",
	EventExpired:                "EVENT_EXPIRED",
	EventConnected:              "EVENT_CONNECTED",
	EventDisconnected:           "EVENT_DISCONNECTED",
	EventReconnected:            "EVENT_RECONNECTED",
	EventLoginSucceeded:         "EVENT_LOGIN_SUCCEEDED",
	EventLoginFailed:            "EVENT_LOGIN_FAILED",
	EventLocationSucceeded:      "EVENT_LOCATION_SUCCEEDED",
	EventLocationFailed:         "EVENT_LOCATION_FAILED",
	EventProtocolViolation:      "EVENT_PROTOCOL_VIOLATION",
	EventUnexpectedResponse:     "EVENT_UNEXPECTED_RESPONSE",
	EventInvalidStateTransition: "EVENT_INVALID_STATE_TRANSITION",
	EventInvalid:                "EVENT_INVALID",
}

func (s SimulatorEventType) String() string {
	if !s.IsValid() {
		return fmt.Sprintf("SIMULATOR_EVENT(%d)", int(s))
	}

	return simulatorEventTypeNames[s]
}

func (s SimulatorEventType) IsValid() bool {
	return s >= 0 && int(s) < len(simulatorEventTypeNames)
}

type SimulatorEvent struct {
	Id   device.Imei
	Type SimulatorEventType
	Time time.Time
}
