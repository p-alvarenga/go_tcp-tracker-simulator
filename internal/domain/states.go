package domain

type DeviceState uint64
type SessionState uint64

const (
	StateDeviceCreated DeviceState = iota
	StateDeviceStopped
	StateDeviceAwaitingLogin
	StateDeviceLocation
	StateDeviceHeartbeat
)

const (
	StateSessionConnected SessionState = iota
	StateSessionReconnecting
	StateSessionClosed
)
