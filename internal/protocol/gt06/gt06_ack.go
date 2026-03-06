package gt06

import (
	"encoding/binary"
	"fmt"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol"
)

func ExtractAck(raw []byte) (*AckPacket, error) {
	if len(raw) != 10 {
		if int(raw[2]) != 5 {
			return nil, fmt.Errorf("gt06: invalid packet length")
		}

		return nil, fmt.Errorf("gt06: invalid packet length")
	}

	packetTypeFlag := raw[3]
	serialNumber := binary.BigEndian.Uint16(raw[4:6])

	if !protocol.CheckCRC(raw) {
		return nil, fmt.Errorf("gt06: invalid crc")
	}

	switch packetTypeFlag {
	case loginFlag:
		return &AckPacket{
			Type:   LoginType,
			Serial: serialNumber,
		}, nil

	case locationFlag:
		return &AckPacket{
			Type:   LocationType,
			Serial: serialNumber,
		}, nil

	default:
		return nil, nil // not valid nor supported protocol type
	}
}
