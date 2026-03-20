package config

import "time"

type SessionConfig struct {
	ServerHost string
	ServerPort int

	DialTimeout time.Duration

	Reconnection *ReconnectConfig
	Lag          *LagConfig
}

type ReconnectConfig struct {
	Enabled           bool
	MaxRetries        uint
	MaxBackOff        time.Duration
	MinBackOff        time.Duration
	BackOffMultiplier float64
}

type LagConfig struct {
	Enabled        bool
	PacketLossRate float64
}

func DefaultSessionConfig() *SessionConfig {
	return &SessionConfig{
		ServerHost: "localhost",
		ServerPort: 9000,

		DialTimeout: 10 * time.Second,

		Reconnection: &ReconnectConfig{
			Enabled:           true,
			MaxRetries:        2,
			MinBackOff:        1 * time.Second,
			MaxBackOff:        5 * time.Minute,
			BackOffMultiplier: 1.2,
		},
	}
}
