package sim

import (
	"context"
	"log/slog"
)

func (s *Simulator) Boot() error {
	s.ctx, s.cancel = context.WithCancel(context.Background())

	err := s.createSimulatedDevices()
	if err != nil {
		s.logger.Error("Error in booting simulator", slog.Any("err", err))
		return err
	}

	s.startSimulatedDevices()
	s.loop()

	return nil
}
