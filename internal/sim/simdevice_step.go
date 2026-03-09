package sim

import (
	"time"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain"
)

func (sd *SimulatedDevice) boot() {
	sd.logger.With("state", sd.getState())

	go sd.ReadSession()
	go sd.MonitorSession()
	go sd.loop()
	go sd.Session.Start(sd.ctx)

	sd.logger.Info("booted device")

	<-sd.ctx.Done()
}

func (sd *SimulatedDevice) loop() {
	ticker := time.NewTicker(sd.cfg.TickInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			st := sd.getState()

			if st != domain.StateDisconnected && st != domain.StateReconnecting {
				sd.step(st)
			}

		case <-sd.ctx.Done():
			return
		}
	}
}

func (sd *SimulatedDevice) step(state domain.SimulatedDeviceState) {
	switch state {
	case domain.StateConnected,
		domain.StateCreated:
		err := sd.SendLogin()
		if err != nil {
			sd.logger.Error("Could not send login packet")
		}

		sd.logger.Info("Sending login")

	case domain.StateLoggedIn:
		sd.logger.Info("Logged in")
	}
}
