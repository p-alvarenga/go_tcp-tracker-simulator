package domain

const (
	EventDeviceConnected DeviceEventType = iota
	EventDeviceRequestedConnection

	EventDeviceLogin
	EventDeviceLocation
	EventDeviceHeartbeat

	EventDeviceError
)

var deviceEventToString = map[DeviceEventType]string{
	EventDeviceConnected:           "EV_DEVICE_CONNECTED",
	EventDeviceRequestedConnection: "EV_DEVICE_REQUESTED_CONNECTION",
	EventDeviceLogin:               "EV_DEVICE_LOGIN",
	EventDeviceLocation:            "EV_DEVICE_LOCATION",
	EventDeviceHeartbeat:           "EV_DEVICE_HEARTBEAT",
	EventDeviceError:               "EV_DEVICE_ERROR",
}

func (e DeviceEventType) String() string {
	if s, ok := deviceEventToString[e]; ok {
		return s
	}

	return "EV_DEVICE_UNKNOWN"
}

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

	EventSessionError
)

var sessionEventToString = map[SessionEventType]string{
	EventSessionConnected:           "EV_SESSION_CONNECTED",
	EventSessionClosed:              "EV_SESSION_CLOSED",
	EventSessionStartedReconnection: "EV_SESSION_STARTED_RECONNECTION",

	EventSessionLoginSucceed:     "EV_SESSION_LOGIN_SUCCEED",
	EventSessionLocationSucceed:  "EV_SESSION_LOCATION_SUCCEED",
	EventSessionHeartbeatSucceed: "EV_SESSION_HEARTBEAT_SUCCEED",

	EventSessionLoginFailed:     "EV_SESSION_LOGIN_FAILED",
	EventSessionLocationFailed:  "EV_SESSION_LOCATION_FAILED",
	EventSessionHeartbeatFailed: "EV_SESSION_HEARTBEAT_FAILED",

	EventSessionError: "EV_SESSION_ERROR",
}

func (e SessionEventType) String() string {
	if s, ok := sessionEventToString[e]; ok {
		return s
	}

	return "EV_SESSION_UNKNOWN"
}
