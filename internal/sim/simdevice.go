package sim

import (
	"context"
	"log/slog"
	"sync"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/config"
	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain"
	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain/device"
	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol/gt06"
	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/tcp"
)

type SimulatedDevice struct {
	simulator *Simulator
	cfg       *config.SimulatedDeviceConfig

	Client *tcp.Client
	Device *device.Device

	reconnecting bool
	ctx          context.Context
	cancel       context.CancelFunc
	mu           sync.Mutex

	logger *slog.Logger

	state      domain.SimulatedDeviceState
	lastPacket gt06.Packet
}

func NewSimulatedDevice(sim *Simulator, client *tcp.Client, device *device.Device, rootLogger *slog.Logger) *SimulatedDevice {
	ctx, cancel := context.WithCancel(sim.ctx)

	logger := rootLogger.With(
		"layer", "SimulatedDevice",
		"imei", device.Imei,
	)

	return &SimulatedDevice{
		simulator: sim,
		cfg:       &sim.cfg.SimulatedDeviceConfig,
		Client:    client,
		Device:    device,
		ctx:       ctx,
		cancel:    cancel,
		logger:    logger,
		state:     domain.StateNew,
	}
}

func (sd *SimulatedDevice) setState(state domain.SimulatedDeviceState) {
	sd.mu.Lock()
	sd.state = state
	sd.mu.Unlock()
}

func (sd *SimulatedDevice) getState() domain.SimulatedDeviceState {
	sd.mu.Lock()
	defer sd.mu.Unlock()
	return sd.state
}

func (sd *SimulatedDevice) Shutdown() {
	sd.cancel()
}
