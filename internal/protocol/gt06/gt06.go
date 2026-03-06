package gt06

type PacketType byte

var (
	startBytes = [2]byte{0x78, 0x78}
	stopBytes  = [2]byte{0x0D, 0x0A}
)

const (
	LoginType    PacketType = 0x01
	LocationType PacketType = 0x12
)

type Packet interface {
	Type() PacketType
	Serial() uint16

	Build() ([]byte, error)
}

func CheckACK(pkt Packet, ack *ACKPacket) bool {
	return pkt.Type() == ack.Type() && pkt.Serial() == ack.Serial()
}
