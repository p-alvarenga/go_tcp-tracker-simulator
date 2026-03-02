package config

import (
	"time"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain"
)

func DefaultConfig() *SimulatorConfig {
	return &SimulatorConfig{
		ServerAddr:  "localhost:9000",
		DialTimeout: 5 * time.Second,
		MaxDevices:  1,

		AutoStart:               true,
		GracefulShutdownTimeout: 10 * time.Second,

		SimulatedDeviceConfig: SimulatedDeviceConfig{
			TickInterval: 1 * time.Second,

			Lag:    defaultLagConfig(),
			Device: defaultDeviceConfig(),
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

		LoginRetry: loginConfig{
			MaxRetries: 5,
			BackoffMin: 1 * time.Second,
			BackoffMax: 5 * time.Second,
		},

		Location: locationConfig{
			Enabled: true,

			Mode:         LocationModeRandomWalk,
			RadiusMeters: 300,

			MaxUpdateInterval: 3 * time.Second,
			MinUpdateInterval: 9 * time.Second,
		},
	}
}
