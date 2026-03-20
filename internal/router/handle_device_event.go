package router

import (
	"go_tcp-tracker-simulator/internal/device"
	"go_tcp-tracker-simulator/internal/domain"
	"go_tcp-tracker-simulator/internal/protocol"
)

func (r *EventRouter) handleDeviceEvent(ev device.DeviceEvent) {
	switch ev.Kind {

	case domain.EventDeviceRequestedConnection:
		id, err := r.sessions.CreateAndRun(ev.IMEI)
		if err != nil {
			r.logger.Error("Could not create session")
		}

		r.logger.Info("Created Session", "id", id)

	case domain.EventDeviceLogin:
		if ev.Packet == nil || ev.Packet.Kind() != protocol.LoginType {
			r.logger.Error("Unexpected error", "err", "EV_DEVICE_LOGIN does not carry a valid Login Packet")
			break
		}

		err := r.sessions.SendPacketIntoSession(ev.SessionID, ev.Packet)
		if err != nil {
			r.logger.Error("Sessions Manager could not send packet", "err", err)
			break
		}

		r.logger.Info("Sent packet", "pkt", ev.Packet)
	}
}
