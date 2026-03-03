package sim

import (
	"fmt"
	"time"
)

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
			sd.logger.Info("Reconnected Successfully", "state", sd.state)
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
