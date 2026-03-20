package device

import (
	"context"
	"go_tcp-tracker-simulator/internal/config"
	"go_tcp-tracker-simulator/internal/domain"
	"go_tcp-tracker-simulator/internal/protocol"
	"log/slog"
	"sync"
)

type DeviceManager struct {
	devices map[domain.IMEI]*Device
	cfg     *config.DeviceConfig

	imeiGenerator *protocol.IMEIGenerator
	sink          chan<- DeviceEvent

	mu sync.RWMutex

	ctx    context.Context
	cancel context.CancelFunc

	rootLogger *slog.Logger
}

func NewDeviceManager(cfg *config.DeviceConfig, sink chan<- DeviceEvent, parentCtx context.Context, rootLogger *slog.Logger) *DeviceManager {
	ctx, cancel := context.WithCancel(parentCtx)

	return &DeviceManager{
		devices: make(map[domain.IMEI]*Device),
		cfg:     cfg,

		sink:          sink,
		imeiGenerator: protocol.NewIMEIGenerator(cfg.IMEI.TACBase, cfg.IMEI.SerialStart),

		ctx:    ctx,
		cancel: cancel,

		rootLogger: rootLogger,
	}
}

func (m *DeviceManager) Get(imei domain.IMEI) *Device {
	m.mu.RLock()
	d, ok := m.devices[imei]
	m.mu.RUnlock()

	if ok {
		return d
	}

	return nil
}

func (m *DeviceManager) Create(imei domain.IMEI) error {
	d := newDevice(imei, 0, m.sink, m.cfg, m.rootLogger)

	m.mu.Lock()
	m.devices[imei] = d
	m.mu.Unlock()

	return nil
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

	d = newDevice(imei, sessionID, m.sink, m.cfg, m.rootLogger)

	m.mu.Lock()
	m.devices[imei] = d
	m.mu.Unlock()

	return d
}

func (m *DeviceManager) CreateDevices(n uint) error {
	for range n {
		imei := m.imeiGenerator.Next()
		m.Create(imei)
	}

	return nil
}
