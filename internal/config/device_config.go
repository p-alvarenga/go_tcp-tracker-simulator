package config

import "time"

type DeviceConfig struct {
	TickInterval time.Duration

	IMEI      *IMEIConfig
	Telemetry *TelemetryConfig
	Protocol  *ProtocolConfig
}

type IMEIConfig struct {
	TACBase     string // 8 bytes
	SerialStart uint   // 6 bytes
}

type TelemetryConfig struct {
	BaseInterval  time.Duration
	JitterPercent float64

	Location LocationConfig
	Network  NetworkConfig
}

type LocationConfig struct{}
type NetworkConfig struct{}

type ProtocolConfig struct {
	PacketSerialStart uint16
}

func DefaultDeviceConfig() *DeviceConfig {
	return &DeviceConfig{
		TickInterval: 1 * time.Second,

		Protocol: &ProtocolConfig{
			PacketSerialStart: 0,
		},

		IMEI: &IMEIConfig{
			TACBase:     "12345678",
			SerialStart: 0,
		},
		Telemetry: &TelemetryConfig{},
	}
}
