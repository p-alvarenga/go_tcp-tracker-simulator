package sim

import (
	"time"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain"
	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol/gt06"
)

func (sd *SimulatedDevice) readClient() {
	for {
		select {
		case <-sd.ctx.Done():
			return

		case ack, ok := <-sd.Client.ReadCh:
			if !ok {
				sd.logger.Warn("Client read channel was closed. shutting simulated device down")

				sd.emit(domain.EventDisconnected)
				return
			}

			ev := sd.translateAck(ack)
			// sd.logger.Info("Ack Event", "event", ev)

			sd.emit(ev)
		}
	}
}

func (sd *SimulatedDevice) translateAck(raw []byte) domain.SimulatorEventType {
	ack, err := gt06.ExtractAck(raw)

	if err != nil || ack == nil {
		sd.logger.Error("Protocol Violation", "err", err)
		return domain.EventProtocolViolation
	}

	if sd.lastPacket == nil {
		return domain.EventUnexpectedResponse
	}

	if !sd.lastPacket.ReceiveAck(ack) {
		return domain.EventUnexpectedResponse
	}

	switch sd.lastPacket.(type) {
	case *gt06.LoginPacket:
		return domain.EventLoginSucceeded

	case *gt06.LocationPacket:
		return domain.EventLocationSucceeded
	}

	return domain.EventInvalid
}

func (sd *SimulatedDevice) emit(eventType domain.SimulatorEventType) {
	sd.simulator.emit(domain.SimulatorEvent{
		Id:   sd.Device.Imei,
		Type: eventType,
		Time: time.Now(),
	})
}
