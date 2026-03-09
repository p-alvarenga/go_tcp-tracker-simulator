package device

import (
	"context"
	"go_tcp-tracker-simulator/internal/domain"
	"log/slog"
	"sync"
)

type DeviceManager struct {
	devices map[domain.IMEI]*Device

	mu sync.RWMutex

	ctx    context.Context
	cancel context.CancelFunc

	rootLogger *slog.Logger
}

func NewDeviceManager(parentCtx context.Context, rootLogger *slog.Logger) *DeviceManager {
	ctx, cancel := context.WithCancel(parentCtx)

	return &DeviceManager{
		devices: make(map[domain.IMEI]*Device),

		ctx:    ctx,
		cancel: cancel,

		rootLogger: rootLogger,
	}
}

func (m *DeviceManager) GetOrCreate(imei domain.IMEI, sessionID domain.SessionID) *Device {
	m.mu.RLock()
	d, ok := m.devices[imei] // first check
	m.mu.Unlock()

	if ok {
		return d
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	d, ok = m.devices[imei]
	if ok {
		return d
	}

	d = newDevice(imei, sessionID, m.ctx, m.rootLogger)

	m.mu.Lock()
	m.devices[imei] = d
	m.mu.Unlock()

	return d
}

func (m *DeviceManager) RunAll() error {
	for _, d := range m.devices {
		d.run()
	}
}
