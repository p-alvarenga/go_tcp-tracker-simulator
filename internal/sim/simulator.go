package sim

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/device"
)

type SimulatorEventType int

const (
	EvDeviceCreated SimulatorEventType = iota
	EvDeviceStarted
	EvDeviceStopped
	EvDeviceDisconnected
	EvDeviceReconnected

	EvDeviceLoginSucceeded
	EvDeviceLocationSucceeded
	EvDeviceLoginFailed
	EvDeviceSessionExpired

	EvDeviceProtocolViolation
	EvDeviceInvalidStateTransition
	EvDeviceUnexpectedResponse

	EvUnknown
)

type SimulatorEvent struct {
	DeviceId device.Imei
	Type     SimulatorEventType
	Time     time.Time
}

type Simulator struct {
	cfg              SimulatorConfig
	simulatedDevices map[device.Imei]*simulatedDevice

	eventCh chan SimulatorEvent

	ctx    context.Context
	cancel context.CancelFunc

	wg sync.WaitGroup
	mu sync.Mutex

	logger *slog.Logger
}

func NewSimulator(cfg *SimulatorConfig) *Simulator {
	return &Simulator{
		cfg:              *cfg,
		simulatedDevices: make(map[device.Imei]*simulatedDevice),
		eventCh:          make(chan SimulatorEvent, 4096),
		logger:           slog.With(slog.String("layer", "sim.Simulator")),
	}
}

func (s *Simulator) registerSimulatedDevice(sd *simulatedDevice) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.simulatedDevices[sd.device.Imei] = sd
}
