package device

import (
	"go_tcp-tracker-simulator/internal/config"
	"go_tcp-tracker-simulator/internal/domain"
	"go_tcp-tracker-simulator/internal/protocol"
	"hash/fnv"
	"log/slog"
	"math/rand"
	"time"
)

type Device struct {
	IMEI      domain.IMEI
	SessionID domain.SessionID
	State     domain.DeviceState

	cfg *config.DeviceConfig

	sink chan<- DeviceEvent

	rng      *rand.Rand
	nextTick time.Time

	logger *slog.Logger
}

func newDevice(imei domain.IMEI, sessionID domain.SessionID, sink chan<- DeviceEvent, cfg *config.DeviceConfig, rootLogger *slog.Logger) *Device {
	return &Device{
		IMEI:      imei,
		SessionID: sessionID,
		cfg:       cfg,

		State: domain.StateDeviceCreated,
		sink:  sink,

		rng: rand.New(rand.NewSource(int64(hashIMEI(imei)))),

		logger: rootLogger.With("lyr", "device"),
	}
}

func (d *Device) emit(kind domain.DeviceEventType, pkt protocol.Packet) {
	d.sink <- DeviceEvent{
		Kind:      kind,
		IMEI:      d.IMEI,
		SessionID: d.SessionID,
		Packet:    pkt,
		Time:      time.Now(),
	}
}

func (d *Device) emitError(err error) {
	d.sink <- DeviceEvent{
		Kind:      domain.EventDeviceError,
		IMEI:      d.IMEI,
		SessionID: d.SessionID,
		Packet:    nil,
		Error:     err,
		Time:      time.Now(),
	}
}

func (d *Device) setState(st domain.DeviceState) { d.State = st }
func hashIMEI(imei domain.IMEI) int64 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(imei))
	return int64(h.Sum32())
}
