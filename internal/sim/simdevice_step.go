package sim

import (
	"fmt"
	"time"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain"
)

func (sd *SimulatedDevice) boot() {
	sd.logger.With("state", sd.getState())

	go sd.ReadClient()
	go sd.MonitorClient()
	go sd.loop()
	go sd.Client.Start(sd.ctx)

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

func (sd *SimulatedDevice) reconnectLoop() error {
	attempt := 0
	backoff := sd.cfg.Connection.BackoffMin

	for {
		select {
		case <-sd.ctx.Done():
			return sd.ctx.Err()
		default:
		}

		if sd.cfg.Connection.MaxRetries > 0 && attempt >= sd.cfg.Connection.MaxRetries {
			return fmt.Errorf("max reconnect attempts reached")
		}

		attempt++

		err := sd.Client.TryConnect()
		if err == nil {
			sd.logger.Info("Reconnected Successfully")
			return nil
		}

		sd.logger.Warn(
			"Could not reconnect",
			"attempt", attempt,
			"backoff", backoff,
			"err", err,
		)

		timer := time.NewTimer(backoff)
		select {
		case <-sd.ctx.Done():
			timer.Stop()
			return sd.ctx.Err()
		case <-timer.C:
		}

		backoff = time.Duration(float64(backoff) * sd.cfg.Connection.BackoffMultiplier)
		if backoff > sd.cfg.Connection.BackoffMax {
			backoff = sd.cfg.Connection.BackoffMax
		}
	}
}
