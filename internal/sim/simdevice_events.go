package sim

import (
	"time"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol/gt06"
)

type SimulatedDeviceEventType int
type AckFailureReason int

const (
	EventAckReceived SimulatedDeviceEventType = iota
	EventAckInvalid
	EventAckWithoutPacket
	EventDisconnected
)

const (
	AckParseFailed AckFailureReason = iota
	AckNoPendingPacket
	AckMismatch
)

type simulatedDeviceEvent struct {
	Type   SimulatedDeviceEventType
	Reason AckFailureReason
}

func (sd *simulatedDevice) listenClient() {
	for {
		select {
		case <-sd.ctx.Done():
			return

		case raw, ok := <-sd.client.AckCh:
			if !ok {
				sd.emit(&SimulatorEvent{
					DeviceId: sd.device.Imei,
					Type:     EvDeviceProtocolViolation,
					Time:     time.Now(),
				})
				return
			}

			ev := sd.translateAck(raw)
			if ev != nil {
				sd.emit(ev)
			}
		}
	}
}

func (sd *simulatedDevice) translateAck(raw []byte) *SimulatorEvent {
	ack, ok := gt06.ExtractAck(raw)

	var ev SimulatorEvent
	ev.DeviceId = sd.device.Imei
	ev.Time = time.Now()
	ev.Type = EvUnknown

	if !ok {
		ev.Type = EvDeviceProtocolViolation
		return &ev
	}

	if sd.lastPacket == nil {
		ev.Type = EvDeviceUnexpectedResponse
		return &ev
	}

	if !sd.lastPacket.ReceiveAck(ack) {
		ev.Type = EvDeviceProtocolViolation
		return &ev
	}

	switch sd.lastPacket.(type) {
	case *gt06.LoginPacket:
		ev.Type = EvDeviceLoginSucceeded
		return &ev

	case *gt06.LocationPacket:
		ev.Type = EvDeviceLocationSucceeded
	}

	return &ev
}

func (sd *simulatedDevice) emit(ev *SimulatorEvent) {
	sd.simulator.emit(ev)
}
