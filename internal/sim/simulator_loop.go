package sim

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

func (s *Simulator) emit(ev *SimulatorEvent) {
	select {
	case s.eventCh <- *ev:
	case <-s.ctx.Done():
		return
	}
}

func (s *Simulator) handleEvent(ev SimulatorEvent) {
	sd := s.simulatedDevices[ev.DeviceId]
	if sd == nil {
		return
	}

	switch ev.Type {
	case EvDeviceProtocolViolation, EvUnknown, EvDeviceUnexpectedResponse:
		// sd.setState()
		s.shutdownSimulatedDevice(ev.DeviceId)
	}
}
