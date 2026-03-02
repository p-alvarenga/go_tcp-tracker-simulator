package sim

import (
	"time"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain"
)

func (sd *SimulatedDevice) boot() {
	go sd.listenClient()
	go sd.loop()
	go sd.Client.Start(sd.ctx)

	sd.logger.Info("Simulated Device Initialized")

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
			sd.step()
		}
	}
}

func (sd *SimulatedDevice) step() {
	switch sd.state {
	case domain.StateNew:
		err := sd.SendLogin()
		if err != nil {
			sd.logger.Error("Could not send login packet")
		}

	case domain.StateLoggedIn:
	}
}
