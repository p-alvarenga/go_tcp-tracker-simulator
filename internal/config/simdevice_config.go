package config

import "time"

type SimulatedDeviceConfig struct {
	TickInterval time.Duration

	Lag    *lagConfig
	Device *deviceConfig
}
