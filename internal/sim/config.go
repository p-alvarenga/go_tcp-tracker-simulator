package sim

import "time"

type SimulatorConfig struct {
	ServerAddr  string
	DialTimeout time.Duration
	MaxDevices  int

	AutoStart               bool
	GracefulShutdownTimeout time.Duration

	TickInterval   time.Duration
	TimeMultiplier float64

	Lag    *LagConfig
	Device *DeviceConfig
}

type LagConfig struct {
	Enabled bool

	Min    time.Duration
	Max    time.Duration
	Jitter time.Duration

	PacketLossRate float64 // [0.0, 1.0]
}

type simulatedDeviceConfig struct {
	TickInterval time.Duration

	Lag    *LagConfig
	Device *DeviceConfig
}

type DeviceConfig struct {
	ImeiTacBase     string
	ImeiSerialStart int

	InitialState SimulatedDeviceState
	LoginRetry   LoginRetryConfig
	Location     LocationConfig
}

type LoginRetryConfig struct {
	MaxRetries int
	BackoffMin time.Duration
	BackoffMax time.Duration
}

type LocationConfig struct {
	Enabled bool

	Mode LocationMode

	RadiusMeters float64

	MaxUpdateInterval time.Duration
	MinUpdateInterval time.Duration
}

type LocationMode int

const (
	LocationModeStatic LocationMode = iota
	LocationModeRandomWalk
	LocationModeRoute
)

func DefaultConfig() *SimulatorConfig {
	return &SimulatorConfig{
		ServerAddr:  "localhost:9000",
		DialTimeout: 5 * time.Second,
		MaxDevices:  1,

		AutoStart:               true,
		GracefulShutdownTimeout: 10 * time.Second,

		TickInterval: 1 * time.Second,

		Lag:    defaultLagConfig(),
		Device: defaultDeviceConfig(),
	}
}

func defaultLagConfig() *LagConfig {
	return &LagConfig{
		Enabled: false,

		Min:    50 * time.Millisecond,
		Max:    300 * time.Millisecond,
		Jitter: 50 * time.Millisecond,

		PacketLossRate: 0.0,
	}
}

func defaultDeviceConfig() *DeviceConfig {
	return &DeviceConfig{
		InitialState: StNew,

		ImeiTacBase:     "12345678",
		ImeiSerialStart: 1,

		LoginRetry: LoginRetryConfig{
			MaxRetries: 5,
			BackoffMin: 1 * time.Second,
			BackoffMax: 5 * time.Second,
		},

		Location: LocationConfig{
			Enabled: true,

			Mode: LocationModeRandomWalk,

			RadiusMeters:      300,
			MaxUpdateInterval: 3 * time.Second,
			MinUpdateInterval: 9 * time.Second,
		},
	}
}
