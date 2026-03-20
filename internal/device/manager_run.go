package device

import "time"

func (m *DeviceManager) Run() error {
	ticker := time.NewTicker(m.cfg.TickInterval)
	defer ticker.Stop()

	for {
		select {
		case <-m.ctx.Done():
			return nil // should it return error?

		case now := <-ticker.C:
			m.tick(now)
		}
	}
}

func (m *DeviceManager) tick(now time.Time) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, d := range m.devices {
		d.tick(now)
	}
}
