package device

import (
	"context"
	"go_tcp-tracker-simulator/internal/domain"
	"log/slog"
)

type Device struct {
	IMEI      domain.IMEI
	SessionID domain.SessionID
	State     domain.DeviceState

	ctx    context.Context
	cancel context.CancelFunc

	logger *slog.Logger

	// Simulation Config
	// Telemetry *Telemetry
}

func newDevice(imei domain.IMEI, sessionID domain.SessionID, parentCtx context.Context, rootLogger *slog.Logger) *Device {
	ctx, cancel := context.WithCancel(parentCtx)

	return &Device{
		IMEI:      imei,
		SessionID: sessionID,
		State:     domain.StateDeviceCreated,

		ctx:    ctx,
		cancel: cancel,

		logger: rootLogger.With("lyr", "device"),
	}
}

func (d *Device) run() {

}
