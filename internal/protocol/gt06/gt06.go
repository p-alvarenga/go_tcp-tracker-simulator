package gt06

var (
	StartFlag []byte = []byte{0x78, 0x78}
	StopFlag  []byte = []byte{0x0D, 0x0A}
)

const (
	LoginFlag    byte = 0x01
	LocationFlag byte = 0x12
)

type PacketType int

const (
	LoginType PacketType = iota
	LocationType
)

type Packet interface {
	Build() ([]byte, error)
	ReceiveAck(*AckPacket) bool
}

type AckPacket struct {
	PacketType   PacketType
	SerialNumber uint16
}

type LoginPacket struct {
	Imei   string
	Serial uint16
}

type LocationPacket struct {
	Serial uint16
}

func (lp *LoginPacket) ReceiveAck(ack *AckPacket) bool {
	return ack.PacketType == LoginType && ack.SerialNumber == lp.Serial
}

func (lp *LocationPacket) ReceiveAck(ack *AckPacket) bool {
	return ack.PacketType == LocationType && ack.SerialNumber == lp.Serial
}
