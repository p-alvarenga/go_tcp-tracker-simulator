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
	select {
	case s.eventCh <- ev:
	case <-s.ctx.Done():
		return
	}
}

func (s *Simulator) handleEvent(event domain.SimulatorEvent) {
	sd := s.simulatedDevices[event.Id]
	if sd == nil {
		return
	}

	switch event.Type {
	case domain.EventProtocolViolation, // probably
		domain.EventUnexpectedResponse,
		domain.EventDisconnected,
		domain.EventUnknown:

		s.shutdownSimulatedDevice(event.Id)
	}
}
