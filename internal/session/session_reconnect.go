package session

import (
	"math"
	"time"
)

func (s *Session) waitRetry(delay time.Duration) bool {
	timer := time.NewTimer(delay)
	select {
	case <-s.ctx.Done():
		timer.Stop()
		return false

	case <-timer.C:
		return true
	}
}

func (s *Session) nextRetryDelay() time.Duration {
	if s.retryCounter == 0 {
		return s.cfg.Reconnection.MinBackOff
	}

	backoff := float64(s.cfg.Reconnection.MinBackOff) * math.Pow(s.cfg.Reconnection.BackOffMultiplier, float64(s.retryCounter))
	d := time.Duration(backoff)

	if d > s.cfg.Reconnection.MaxBackOff {
		return s.cfg.Reconnection.MaxBackOff
	}

	return d
}
