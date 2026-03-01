package sim

import (
	"log/slog"
	"net"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/device"
	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol"
	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/tcp"
)

func (s *Simulator) createSimulatedDevices() error {
	imeiGenerator := protocol.NewImeiGenerator(s.cfg.Device.ImeiTacBase, s.cfg.Device.ImeiSerialStart)
	deviceConfig := &simulatedDeviceConfig{
		TickInterval: s.cfg.TickInterval,

		Lag:    s.cfg.Lag,
		Device: s.cfg.Device,
	}

	for range s.cfg.MaxDevices {
		imei := imeiGenerator.Next()

		conn, err := net.DialTimeout("tcp", s.cfg.ServerAddr, s.cfg.DialTimeout)
		if err != nil {
			s.logger.Error("Could not connect into server", slog.String("addr", s.cfg.ServerAddr))
		}

		client := tcp.NewClient(conn)
		device := device.NewDevice(imei)

		sd := newSimulatedDevice(s, client, device, deviceConfig)
		s.registerSimulatedDevice(sd)
	}

	return nil
}

func (s *Simulator) startSimulatedDevices() {
	for _, sd := range s.simulatedDevices {
		go sd.run()
	}
}

func (s *Simulator) shutdownSimulatedDevice(sdId device.Imei) {
	sd := s.simulatedDevices[sdId]
	if sd == nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	sd.Shutdown()
	delete(s.simulatedDevices, sdId)
}
