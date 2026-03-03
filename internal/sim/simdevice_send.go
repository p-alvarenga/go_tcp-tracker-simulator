package sim

import "github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol/gt06"

func (sd *SimulatedDevice) SendLogin() error {
	pkt, err := gt06.NewLoginPacket(string(sd.Device.Imei))

	if err != nil {
		return err
	}

	err = sd.Client.SendPacket(pkt)
	if err != nil {
		sd.logger.Error("Could not send packet", "err", err)
	}

	sd.lastPacket = pkt

	return nil
}
