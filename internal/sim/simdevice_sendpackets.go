package sim

import "gt06_sim/internal/protocol/gt06"

func (sd *simulatedDevice) SendLogin() error {
	pkt, err := gt06.NewLoginPacket(string(sd.device.Imei))

	if err != nil {
		return err
	}

	sd.client.SendPacket(pkt)

	return nil
}
