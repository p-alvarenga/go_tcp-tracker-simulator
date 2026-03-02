package gt06

import (
	"encoding/binary"
	"fmt"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol"
)

func NewLoginPacket(imei string) (*LoginPacket, error) {
	if !protocol.ValidateImei(imei) {
		return nil, fmt.Errorf("gt06: Invalid imei string")
	}

	return &LoginPacket{
		Imei:   imei,
		Serial: 0,
	}, nil
}

func (lp *LoginPacket) Type() PacketType {
	return LoginType
}

func (lp *LoginPacket) Build() ([]byte, error) {
	raw := make([]byte, 0, 32)

	raw = append(raw, startFlag[:]...)
	raw = append(raw, []byte{
		0x0D,
		loginFlag,
	}...)

	bcdImei, err := protocol.EncodeImeiToBcd(lp.Imei)
	if err != nil {
		return nil, fmt.Errorf("gt06: Could not encode imei \"%s\" to bcd", lp.Imei)
	}

	raw = append(raw, bcdImei...) // payload

	raw = binary.BigEndian.AppendUint16(raw, lp.Serial)
	raw = binary.BigEndian.AppendUint16(raw, protocol.CalculateCRC(raw[2:]))

	raw = append(raw, stopFlag[:]...)

	return raw, nil
}
