package router

import (
	"fmt"
	"go_tcp-tracker-simulator/internal/domain"
	"go_tcp-tracker-simulator/internal/session"
)

func (r *EventRouter) handleSessionEvent(ev session.SessionEvent) {
	switch ev.Kind {
	case domain.EventSessionConnected:

		d := r.devices.Get(ev.IMEI)
		if d == nil {
			r.logger.Error("Could not find device. Fatal error", "imei", ev.IMEI)
			break
		}

		d.State = domain.StateDeviceConnected
		d.SessionID = ev.SessionID

		r.logger.Info(fmt.Sprintf("Device %s connected", ev.IMEI))
	}
}
