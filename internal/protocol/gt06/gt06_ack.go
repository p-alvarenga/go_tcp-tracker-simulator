package gt06

import (
	"encoding/binary"
	"gt06_sim/internal/protocol"
)

func ExtractAck(raw []byte) (*AckPacket, bool) {
	length := int(raw[2])
	if length != 5 || len(raw) != 10 {
		return nil, false
	}

	protocolFlag := raw[3]
	serialNumber := binary.BigEndian.Uint16(raw[4:6])
	crc := binary.BigEndian.Uint16(raw[6:8])
	expectedCrc := protocol.CalculateCRC(raw[2:6])

	if crc != expectedCrc {
		return nil, false
	}

	switch protocolFlag {
	case LoginFlag:
		return &AckPacket{
			PacketType:   LoginType,
			SerialNumber: serialNumber,
		}, true

	case LocationFlag:
		return &AckPacket{
			PacketType:   LocationType,
			SerialNumber: serialNumber,
		}, true

	default:
		return nil, false // not valid nor supported protocol type
	}
}
