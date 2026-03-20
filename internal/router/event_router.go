package router

import (
	"context"
	"fmt"
	"go_tcp-tracker-simulator/internal/config"
	"go_tcp-tracker-simulator/internal/device"
	"go_tcp-tracker-simulator/internal/session"
	"log/slog"
)

type EventRouter struct {
	devices  *device.DeviceManager
	sessions *session.SessionManager

	deviceEvents  chan device.DeviceEvent
	sessionEvents chan session.SessionEvent

	ctx    context.Context
	cancel context.CancelFunc

	logger     *slog.Logger
	rootLogger *slog.Logger
}

func NewEventRouter(sessionsCfg *config.SessionConfig, devicesCfg *config.DeviceConfig, parentCtx context.Context, rootLogger *slog.Logger) *EventRouter {
	ctx, cancel := context.WithCancel(parentCtx)

	r := &EventRouter{
		deviceEvents:  make(chan device.DeviceEvent, 64), // low buffer => controlled backpressure
		sessionEvents: make(chan session.SessionEvent, 64),

		ctx:    ctx,
		cancel: cancel,

		logger:     rootLogger.With("lyr", "EventRouter"),
		rootLogger: rootLogger,
	}

	r.devices = device.NewDeviceManager(devicesCfg, r.deviceEvents, r.ctx, r.rootLogger)
	r.sessions = session.NewSessionManager(sessionsCfg, r.sessionEvents, r.ctx, r.rootLogger)

	return r
}

func (r *EventRouter) Boot(n uint) error {
	err := r.devices.CreateDevices(n)
	if err != nil {
		return err
	}

	r.logger.Info(fmt.Sprintf("Device manager created %d devices", n))

	return nil
}

func (r *EventRouter) Start() error {
	errors := make(chan error, 2)

	r.logger.Info("Router started running")

	go func() { errors <- r.devices.Run() }()

	for {
		select {
		case deviceEvent := <-r.deviceEvents:
			r.handleDeviceEvent(deviceEvent)

		case sessionEvent := <-r.sessionEvents:
			r.handleSessionEvent(sessionEvent)

		case <-r.ctx.Done():
			return r.ctx.Err()

		case err := <-errors:
			r.logger.Error("Event router unexpected error", "err", err)
		}
	}
}
