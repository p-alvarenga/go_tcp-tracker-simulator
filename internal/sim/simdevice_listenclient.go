package sim

import (
	"time"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain"
	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol/gt06"
)

func (sd *SimulatedDevice) listenClient() {
	for {
		select {
		case <-sd.ctx.Done():
			return

		case ack, ok := <-sd.Client.AckCh:
			sd.logger.Info("Received ACK", "ack", ack)

			if !ok {
				sd.emit(domain.EventDisconnected)
				return
			}

			sd.emit(sd.translateAck(ack))
		}
	}
}

func (sd *SimulatedDevice) translateAck(raw []byte) domain.SimulatorEventType {
	if sd.lastPacket == nil {
		return domain.EventUnexpectedResponse
	}

	ack, ok := gt06.ExtractAck(raw)

	if !ok || !sd.lastPacket.ReceiveAck(ack) {
		return domain.EventProtocolViolation
	}

	switch sd.lastPacket.(type) {
	case *gt06.LoginPacket:
		return domain.EventLoginSucceeded

	case *gt06.LocationPacket:
		return domain.EventLocationSucceeded

	default:
		return domain.EventUnknown
	}
}

func (sd *SimulatedDevice) emit(eventType domain.SimulatorEventType) {
	sd.simulator.emit(domain.SimulatorEvent{
		Id:   sd.Device.Imei,
		Type: eventType,
		Time: time.Now(),
	})
}
