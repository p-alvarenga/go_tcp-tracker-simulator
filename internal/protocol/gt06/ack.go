package gt06

import (
	"encoding/binary"
	"fmt"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol"
)

type ACKPacket struct {
	packetType PacketType
	serial     uint16
}

func DecodeACK(raw []byte) (*ACKPacket, error) {
	if len(raw) != 10 {
		if int(raw[2]) != 0x05 {
			return nil, fmt.Errorf("gt06: invalid packet length")
		}
		return nil, fmt.Errorf("gt06: invalid packet length")
	}

	serial := binary.BigEndian.Uint16(raw[4:6])

	if !protocol.CheckCRC(raw) {
		return nil, fmt.Errorf("gt06: invalid crc")
	}

	switch PacketType(raw[3]) {
	case LoginType:
		return &ACKPacket{
			packetType: LoginType,
			serial:     serial,
		}, nil

	case LocationType:
		return &ACKPacket{
			packetType: LocationType,
			serial:     serial,
		}, nil

	default:
		return nil, fmt.Errorf("gt06: not valid or not supported packet type (%X)", raw[3])
	}
}

func (p *ACKPacket) Type() PacketType {
	return p.packetType
}

func (p *ACKPacket) Serial() uint16 {
	return p.serial
}
