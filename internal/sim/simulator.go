package sim

import (
	"context"
	"log/slog"
	"sync"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/config"
	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain"
	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain/device"
)

type Simulator struct {
	cfg              config.SimulatorConfig
	simulatedDevices map[device.Imei]*SimulatedDevice

	eventCh chan domain.SimulatorEvent

	ctx    context.Context
	cancel context.CancelFunc

	wg sync.WaitGroup
	mu sync.Mutex

	logger *slog.Logger
}

func NewSimulator(config *config.SimulatorConfig) *Simulator {
	return &Simulator{
		cfg:              *config,
		simulatedDevices: make(map[device.Imei]*SimulatedDevice),
		eventCh:          make(chan domain.SimulatorEvent, 4096),
		logger:           slog.With(slog.String("layer", "Simulator")),
	}
}

func (s *Simulator) registerSd(sd *SimulatedDevice) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.simulatedDevices[sd.Device.Imei] = sd
}
