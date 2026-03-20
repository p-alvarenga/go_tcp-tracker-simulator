package device

import (
	"go_tcp-tracker-simulator/internal/domain"
	"go_tcp-tracker-simulator/internal/protocol"
	"time"
)

func (d *Device) tick(now time.Time) error {
	if now.Before(d.nextTick) {
		return nil
	}

	switch d.State {

	case domain.StateDeviceCreated:
		d.emit(
			domain.EventDeviceRequestedConnection,
			nil,
		)

		d.setState(domain.StateDeviceWaitingConnection)

	case domain.StateDeviceConnected:
		d.emit(
			domain.EventDeviceLogin,
			&protocol.LoginPacket{
				IMEI:   d.IMEI,
				Serial: d.cfg.Protocol.PacketSerialStart,
			},
		)
	}

	d.nextTick = now.Add(d.computeNextInterval())

	return nil
}

func (d *Device) computeNextInterval() time.Duration {

	if d.cfg.Telemetry.JitterPercent <= 0 {
		return d.cfg.Telemetry.BaseInterval
	}

	max := float64(d.cfg.Telemetry.BaseInterval) * d.cfg.Telemetry.JitterPercent
	jitter := d.rng.Float64() * max

	return d.cfg.Telemetry.BaseInterval + time.Duration(jitter)
}
