package gt06

import (
	"encoding/binary"
	"fmt"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol"
)

type LoginPacket struct {
	IMEI   string
	serial uint16
}

func NewLoginPacket(imei string) (*LoginPacket, error) {
	if !protocol.CheckIMEI(imei) {
		return nil, fmt.Errorf("gt06: Invalid imei string")
	}

	return &LoginPacket{
		IMEI:   imei,
		serial: 0,
	}, nil
}

func (p *LoginPacket) Build() ([]byte, error) {
	raw := make([]byte, 0, 32)

	raw = append(raw, startBytes[:]...)
	raw = append(raw, []byte{
		0x0D,
		byte(LoginType),
	}...)

	bcdImei, err := protocol.IMEIToBcd(p.IMEI)
	if err != nil {
		return nil, fmt.Errorf("gt06: Could not encode imei \"%s\" to bcd", p.IMEI)
	}

	raw = append(raw, bcdImei...) // payload

	raw = binary.BigEndian.AppendUint16(raw, p.serial)
	raw = binary.BigEndian.AppendUint16(raw, protocol.CalculateCRC(raw[2:]))

	raw = append(raw, stopBytes[:]...)

	return raw, nil
}

func (p LoginPacket) Type() PacketType {
	return LoginType
}

func (p LoginPacket) Serial() uint16 {
	return p.serial
}
