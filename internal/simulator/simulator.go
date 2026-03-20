package simulator

import (
	"context"
	"go_tcp-tracker-simulator/internal/config"
	"go_tcp-tracker-simulator/internal/router"
	"log/slog"
)

type Simulator struct {
	cfg *config.Config

	router *router.EventRouter

	ctx    context.Context
	cancel context.CancelFunc

	logger     *slog.Logger
	rootLogger *slog.Logger
}

func NewSimulator(cfg config.Config, rootLogger *slog.Logger) *Simulator {
	ctx, cancel := context.WithCancel(context.Background())

	return &Simulator{
		cfg: &cfg,

		router: router.NewEventRouter(cfg.SessionConfig, cfg.DeviceConfig, ctx, rootLogger),

		ctx:    ctx,
		cancel: cancel,

		logger:     rootLogger.With("lyr", "Simulator"),
		rootLogger: rootLogger,
	}
}

func (s *Simulator) Boot() {
	err := s.router.Boot(s.cfg.NumberOfDevices)

	if err != nil {
		s.logger.Error("Could not boot", "err", err)
		return
	}

	s.logger.Info("Simulator booting")
	s.Start()
}

func (s *Simulator) Start() {
	errors := make(chan error, 1)

	go func() {
		errors <- s.router.Start()
	}()

	select {
	case <-s.ctx.Done():
		s.logger.Info("Simulator shutting down")
	case err := <-errors:
		s.logger.Error("Router stopped", "err", err)
		s.cancel()
	}
}
