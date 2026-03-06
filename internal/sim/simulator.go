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
	simulatedDevices map[device.IMEI]*SimulatedDevice

	eventCh chan domain.SimulatorEvent

	ctx    context.Context
	cancel context.CancelFunc

	wg sync.WaitGroup
	mu sync.Mutex

	rootLogger *slog.Logger
	logger     *slog.Logger
}

func NewSimulator(config *config.SimulatorConfig, rootLogger *slog.Logger) *Simulator {
	return &Simulator{
		cfg:              *config,
		simulatedDevices: make(map[device.IMEI]*SimulatedDevice),
		eventCh:          make(chan domain.SimulatorEvent, 4096),
		logger:           rootLogger.With(slog.String("layer", "Simulator")),
		rootLogger:       rootLogger,
	}
}

func (s *Simulator) Boot() error {
	s.ctx, s.cancel = context.WithCancel(context.Background())

	s.logger.InfoContext(s.ctx, "Simulator booted")

	err := s.createSimulatedDevices()
	if err != nil {
		s.logger.Error("Error in booting simulator", slog.Any("err", err))
		return err
	}

	s.startSimulatedDevices()

	go s.loop()
	<-s.ctx.Done()

	return nil
}
