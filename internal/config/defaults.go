package config

import (
	"time"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain"
)

func DefaultConfig() *SimulatorConfig {
	return &SimulatorConfig{
		ServerAddr: "localhost:9000",
		MaxDevices: 1,

		AutoStart:               true,
		GracefulShutdownTimeout: 10 * time.Second,

		SimulatedDeviceConfig: SimulatedDeviceConfig{
			TickInterval: 1 * time.Second,

			Connection: defaultConnectionConfig(),
			Lag:        defaultLagConfig(),
			Device:     defaultDeviceConfig(),
		},
	}
}

func defaultLagConfig() *lagConfig {
	return &lagConfig{
		Enabled: false,

		Min: 50 * time.Millisecond,
		Max: 300 * time.Millisecond,

		PacketLossRate: 0.0,
	}
}

func defaultDeviceConfig() *deviceConfig {
	return &deviceConfig{
		InitialState: domain.StateNew,

		ImeiTacBase:     "12345678",
		ImeiSerialStart: 1,

		Location: locationConfig{
			Enabled: true,

			Mode:         LocationModeRandomWalk,
			RadiusMeters: 300,

			MaxUpdateInterval: 3 * time.Second,
			MinUpdateInterval: 9 * time.Second,
		},
	}
}

func defaultConnectionConfig() *ConnectionConfig {
	return &ConnectionConfig{
		// Connection
		DialTimeout: 5 * time.Second,

		// Reconnect
		TryReconnect:      true,
		MaxRetries:        0, // infinite
		BackoffMin:        1 * time.Second,
		BackoffMax:        120 * time.Second,
		BackoffMultiplier: 2.0, // 1s -> 2s -> 4s -> 8s -> 16s -> ...
	}
}
