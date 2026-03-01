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
	var raw []byte

	raw = append(raw, StartFlag...)
	raw = append(raw, []byte{
		0x0D,
		LoginFlag,
	}...)

	bcdImei, err := protocol.EncodeImeiToBcd(lp.Imei)
	if err != nil || len(bcdImei) != 15 {
		return nil, fmt.Errorf("gt06: Could not encode imei \"%s\" to bcd", lp.Imei)
	}

	raw = append(raw, bcdImei...) // payload

	raw = binary.BigEndian.AppendUint16(raw, lp.Serial)
	raw = binary.BigEndian.AppendUint16(raw, protocol.CalculateCRC(raw[2:]))

	raw = append(raw, StopFlag...)

	return raw, nil
}
