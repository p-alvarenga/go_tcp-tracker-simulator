package sim

import (
	"time"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain"
)

func (sd *SimulatedDevice) boot() {
	sd.logger.With("state", sd.getState())

	go sd.ReadClient()
	go sd.MonitorClient()
	go sd.loop()
	go sd.Client.Start(sd.ctx)

	sd.logger.Info("booted device")

	<-sd.ctx.Done()
}

func (sd *SimulatedDevice) loop() {
	ticker := time.NewTicker(sd.cfg.TickInterval)
	defer ticker.Stop()

	for {
		select {
		case <-sd.ctx.Done():
			return
		case <-ticker.C:
			st := sd.getState()

			if st != domain.StateDisconnected && st != domain.StateReconnecting {
				sd.step(st)
			}
		}
	}
}

func (sd *SimulatedDevice) step(state domain.SimulatedDeviceState) {
	switch state {
	case domain.StateNew:
		err := sd.SendLogin()
		if err != nil {
			sd.logger.Error("Could not send login packet")
		}

	case domain.StateLoggedIn:
		sd.logger.Info("Logged in")
	}
}
