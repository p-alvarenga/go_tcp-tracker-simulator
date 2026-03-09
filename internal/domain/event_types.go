package domain

const (
	EventDeviceConnected DeviceEventType = iota
	EventDeviceDisconnected

	EventDeviceLogin
	EventDeviceLocation
	EventDeviceHeartbeat

	EventDeviceError
)

const (
	EventSessionConnected SessionEventType = iota
	EventSessionClosed
	EventSessionStartedReconnection

	EventSessionLoginSucceed
	EventSessionLocationSucceed
	EventSessionHeartbeatSucceed

	EventSessionLoginFailed
	EventSessionLocationFailed
	EventSessionHeartbeatFailed
)
