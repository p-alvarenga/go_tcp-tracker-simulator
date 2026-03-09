package sim

import (
	"time"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain"
	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol/gt06"
)

func (sd *SimulatedDevice) ReadSession() {
	for {
		select {
		case <-sd.ctx.Done():
			return

		case ack, ok := <-sd.Session.ACKCh:
			sd.logger.Info("Reading ACK: ", "ack", ack)

			if !ok {
				sd.logger.Warn("Session read channel was closed")
				sd.emit(domain.EventDisconnected)

				return
			}

			ev := sd.translateACK(ack)
			sd.emit(ev)
		}
	}
}

func (sd *SimulatedDevice) MonitorSession() {
	for {
		select {
		case <-sd.ctx.Done():
			return
		case <-sd.Session.Done():
			st := sd.getState()

			if st != domain.StateReconnecting && st != domain.StateConnected {
				sd.emit(domain.EventStartReconnection)

				err := sd.reconnectLoop()
				if err != nil {
					sd.logger.Error("Could not reconnect")
					continue
				}

				sd.emit(domain.EventReconnected)
			}
		}
	}
}

func (sd *SimulatedDevice) translateACK(ack *gt06.ACKPacket) domain.SimulatorEventType {
	if sd.lastPacket == nil {
		return domain.EventUnexpectedResponse
	}

	err := gt06.CheckACK(ack, sd.lastPacket)
	if err != nil {
		sd.logger.Error("Could not translate ack", "err", err)
		return domain.EventUnexpectedResponse
	}

	switch sd.lastPacket.Type() {
	case gt06.LoginType:
		return domain.EventLoginSucceeded

	case gt06.LocationType:
		return domain.EventLocationSucceeded
	}

	return domain.EventInvalid
}

func (sd *SimulatedDevice) emit(eventType domain.SimulatorEventType) {
	sd.simulator.emit(domain.SimulatorEvent{
		ID:   sd.Device.IMEI,
		Type: eventType,
		Time: time.Now(),
	})
}
