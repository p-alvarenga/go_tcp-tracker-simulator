package gt06

import "fmt"

type PacketType byte

var (
	startBytes = [2]byte{0x78, 0x78}
	stopBytes  = [2]byte{0x0D, 0x0A}
)

const (
	LoginType    PacketType = 0x01
	LocationType PacketType = 0x12
)

func (p PacketType) String() string {
	switch p {
	case LoginType:
		return "LOGIN_PACKET"
	case LocationType:
		return "LOCATION_PACKET"
	default:
		return "INVALID_OR_UNSUPPORTED_PACKET"
	}
}

type Packet interface {
	Type() PacketType
	Serial() uint16

	Build() ([]byte, error)
}

func CheckACK(ack *ACKPacket, pkt Packet) error {
	if ack.Type() != pkt.Type() {
		return fmt.Errorf("ACK type (%s) differs from packet type (%s)", ack.Type(), pkt.Type())
	}

	if ack.Serial() != pkt.Serial() {
		return fmt.Errorf("ACK serial (%d) differs from Packet Serial (%d)", ack.Serial(), pkt.Serial())
	}

	return nil
}
