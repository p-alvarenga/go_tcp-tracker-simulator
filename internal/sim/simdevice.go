package sim

import (
	"context"
	"gt06_sim/internal/device"
	"gt06_sim/internal/protocol/gt06"
	"gt06_sim/internal/tcp"
	"log/slog"
)

type SimulatedDeviceState int

const (
	StNew SimulatedDeviceState = iota
	StLoggedIn
)

type simulatedDevice struct {
	simulator *Simulator
	cfg       *simulatedDeviceConfig

	client *tcp.Client
	device *device.Device

	ctx    context.Context
	cancel context.CancelFunc

	logger *slog.Logger

	state             SimulatedDeviceState
	RetryLoginCounter int
	lastPacket        gt06.Packet
}

func newSimulatedDevice(sim *Simulator, c *tcp.Client, d *device.Device, cfg *simulatedDeviceConfig) *simulatedDevice {
	ctx, cancel := context.WithCancel(sim.ctx)

	logger := slog.Default().With(
		"layer", "sim.simulatedDevice",
		"imei", d.Imei,
	)

	return &simulatedDevice{
		simulator:         sim,
		cfg:               cfg,
		client:            c,
		device:            d,
		ctx:               ctx,
		cancel:            cancel,
		logger:            logger,
		state:             StNew,
		RetryLoginCounter: 0,
	}
}

func (sd *simulatedDevice) Shutdown() {
	sd.cancel()
}
