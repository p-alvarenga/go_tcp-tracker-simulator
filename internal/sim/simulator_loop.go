package sim

import (
	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain"
)

func (s *Simulator) loop() {
	for {
		select {
		case ev := <-s.eventCh:
			s.handleEvent(ev)
		case <-s.ctx.Done():
			return
		}
	}
}

func (s *Simulator) emit(ev domain.SimulatorEvent) {
	if !ev.Type.IsValid() {
		s.logger.Warn("Invalid event type tried to emit", "event", ev)
		return
	}

	s.eventCh <- ev
}

func (s *Simulator) handleEvent(event domain.SimulatorEvent) {
	sd := s.simulatedDevices[event.Id]
	if sd == nil {
		return
	}

	switch event.Type {
	case domain.EventReconnected:
		sd.setState(domain.StateConnected)

	case domain.EventLoginSucceeded:
		sd.logger.Info("Device logged")
		sd.setState(domain.StateLoggedIn)

	case domain.EventDisconnected,
		domain.EventStartReconnection:
		s.logger.Warn("Device Client started reconnection", "state", sd.state)
		sd.setState(domain.StateReconnecting)

	case domain.EventProtocolViolation,
		domain.EventUnexpectedResponse,
		domain.EventInvalid: // probably
		s.logger.Error("Shutting down device", "event", event.Type)
		s.shutdownSimulatedDevice(event.Id)
	}
}
