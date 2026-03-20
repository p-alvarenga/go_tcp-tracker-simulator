package domain

// Device
const (
	StateDeviceCreated DeviceState = iota
	StateDeviceStopped

	StateDeviceWaitingConnection

	StateDeviceConnected
	StateDeviceDisconnected

	StateDeviceNotLogged
	StateDeviceSendLocation  // keep sending location packet
	StateDeviceSendHeartbeat // keep sending heartbeat

)

// Session
const (
	StateSessionCreated SessionState = iota
	StateSessionConnected
	StateSessionDisconnected
	StateSessionReconnecting
	StateSessionClosed
)
