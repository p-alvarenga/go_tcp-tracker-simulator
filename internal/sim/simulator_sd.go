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

	addr := net.JoinHostPort(s.cfg.ServerHost, s.cfg.ServerPort)

	for range s.cfg.MaxDevices {
		imei := imeiGenerator.Next()

		client, err := tcp.NewClient(
			addr,
			s.cfg.SimulatedDeviceConfig.TickInterval,
			s.rootLogger,
		)

		if err != nil {
			s.logger.Error("Could not connect into server", slog.String("addr", addr))
			return err
		}

		device := device.NewDevice(imei, s.cfg.SimulatedDeviceConfig.Device.ImeiSerialStart)
		sd := NewSimulatedDevice(s, client, device, s.rootLogger)

		s.registerSimulatedDevice(sd)
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

func (s *Simulator) registerSimulatedDevice(sd *SimulatedDevice) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.simulatedDevices[sd.Device.Imei] = sd
}
