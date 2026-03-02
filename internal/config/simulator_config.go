package config

import (
	"time"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain"
)

type SimulatorConfig struct {
	ServerAddr  string
	DialTimeout time.Duration
	MaxDevices  int

	AutoStart               bool
	GracefulShutdownTimeout time.Duration

	TimeMultiplier float64

	SimulatedDeviceConfig SimulatedDeviceConfig
}

type lagConfig struct {
	Enabled bool

	Min time.Duration
	Max time.Duration

	PacketLossRate float64
}

type deviceConfig struct {
	ImeiTacBase     string
	ImeiSerialStart int

	InitialState domain.SimulatedDeviceState
	LoginRetry   loginConfig
	Location     locationConfig
}

type loginConfig struct {
	MaxRetries int
	BackoffMin time.Duration
	BackoffMax time.Duration
}

type locationConfig struct {
	Enabled bool

	Mode         locationMode
	RadiusMeters float64

	MaxUpdateInterval time.Duration
	MinUpdateInterval time.Duration
}

type locationMode int

const (
	LocationModeStatic locationMode = iota
	LocationModeRandomWalk
	LocationModeRoute
)
