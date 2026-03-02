package sim

import (
	"log/slog"
	"net"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain/device"
	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol"
	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/tcp"
)

func (s *Simulator) createSimulatedDevices() error {
	imeiGenerator := protocol.NewImeiGenerator(
		s.cfg.SimulatedDeviceConfig.Device.ImeiTacBase,
		s.cfg.SimulatedDeviceConfig.Device.ImeiSerialStart,
	)

	for range s.cfg.MaxDevices {
		imei := imeiGenerator.Next()

		conn, err := net.DialTimeout("tcp", s.cfg.ServerAddr, s.cfg.DialTimeout)
		if err != nil {
			s.logger.Error("Could not connect into server", slog.String("addr", s.cfg.ServerAddr))
		}

		client := tcp.NewClient(conn)
		device := device.NewDevice(imei)

		sd := NewSimulatedDevice(s, client, device, &s.cfg.SimulatedDeviceConfig)
		s.registerSd(sd)
	}

	return nil
}

func (s *Simulator) startSimulatedDevices() {
	for _, sd := range s.simulatedDevices {
		go sd.boot()
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
