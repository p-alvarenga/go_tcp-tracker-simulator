package config

type Config struct {
	NumberOfDevices uint

	DeviceConfig  *DeviceConfig
	SessionConfig *SessionConfig
}

func DefaultConfig() *Config {
	return &Config{
		NumberOfDevices: 1,

		DeviceConfig:  DefaultDeviceConfig(),
		SessionConfig: DefaultSessionConfig(),
	}
}
