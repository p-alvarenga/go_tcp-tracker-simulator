package sim

import "time"

func (sd *simulatedDevice) run() {
	for {
		select {
		case <-sd.ctx.Done():
			return
		default:
		}

		go sd.listenClient()
		go sd.startSimulation()
		go sd.loop()
	}
}

func (sd *simulatedDevice) loop() {
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

func (sd *simulatedDevice) step() {
	switch sd.state {
	case StNew:
		_ = sd.SendLogin()
	}
}
