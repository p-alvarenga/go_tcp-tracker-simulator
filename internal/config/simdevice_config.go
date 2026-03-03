package config

import "time"

type SimulatedDeviceConfig struct {
	TickInterval time.Duration

	Lag    *lagConfig
	Device *deviceConfig

	Connection *ConnectionConfig
}

type ConnectionConfig struct {
	DialTimeout time.Duration

	TryReconnect      bool
	MaxRetries        int
	BackoffMin        time.Duration
	BackoffMax        time.Duration
	BackoffMultiplier float64
}
