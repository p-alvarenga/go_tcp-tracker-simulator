package protocol

import (
	"encoding/binary"
	"go_tcp-tracker-simulator/internal/domain"
)

type LoginPacket struct {
	IMEI   domain.IMEI
	Serial uint16
}

func (p *LoginPacket) Build() ([]byte, error) {
	buf := make([]byte, 0, 18)

	bcdImei, err := bcdEncodeIMEI(string(p.IMEI))
	if err != nil {
		return nil, err
	}

	buf = append(buf, startBytes...)
	buf = append(buf, []byte{
		0x0D, // lengh
		byte(LoginType),
	}...)
	buf = append(buf, bcdImei...)
	buf = binary.BigEndian.AppendUint16(buf, p.Serial)
	buf = binary.BigEndian.AppendUint16(buf, CalculateCRC(buf[3:]))
	buf = append(buf, stopBytes...)

	return buf, nil
}

func (p *LoginPacket) IsValid() bool     { return true }
func (p *LoginPacket) GetSerial() uint16 { return p.Serial }
func (p *LoginPacket) Kind() PacketType  { return LoginType }
