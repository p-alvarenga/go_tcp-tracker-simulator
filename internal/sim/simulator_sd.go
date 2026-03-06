package sim

import (
	"log/slog"
	"net"
	"strconv"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain/device"
	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol"
	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/session"
)

func (s *Simulator) createSimulatedDevices() error {
	imeiGenerator := protocol.NewIMEIGenerator(
		s.cfg.SimulatedDeviceConfig.Device.IMEITacBase,
		s.cfg.SimulatedDeviceConfig.Device.IMEISerialStart,
	)

	addr := net.JoinHostPort(s.cfg.ServerHost, strconv.Itoa(s.cfg.ServerPort))

	for range s.cfg.MaxDevices {
		imei := imeiGenerator.Next()

		session, err := session.New(
			addr,
			s.cfg.SimulatedDeviceConfig.TickInterval,
			s.rootLogger,
		)

		if err != nil {
			s.logger.Error("Could not connect into server", slog.String("addr", addr))
			return err
		}

		device := device.NewDevice(imei, s.cfg.SimulatedDeviceConfig.Device.IMEISerialStart)
		sd := NewSimulatedDevice(s, session, device, s.rootLogger)

		s.registerSimulatedDevice(sd)
	}

	return nil
}

func (s *Simulator) startSimulatedDevices() {
	for _, sd := range s.simulatedDevices {
		go sd.boot()
	}
}

func (s *Simulator) shutdownSimulatedDevice(id device.IMEI) { // Double check for race
	s.mu.Lock()
	defer s.mu.Unlock()

	sd := s.simulatedDevices[id]
	if sd == nil {
		return
	}

	sd.Shutdown()
	delete(s.simulatedDevices, id)
}

func (s *Simulator) registerSimulatedDevice(sd *SimulatedDevice) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.simulatedDevices[sd.Device.IMEI] = sd
}
