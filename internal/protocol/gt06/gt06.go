package gt06

type PacketType int

var (
	startFlag = [2]byte{0x78, 0x78}
	stopFlag  = [2]byte{0x0D, 0x0A}
)

const (
	loginFlag    byte = 0x01
	locationFlag byte = 0x12
)

const (
	LoginType PacketType = iota
	LocationType
)

type Packet interface {
	Build() ([]byte, error)
	ReceiveAck(*AckPacket) bool
}

type AckPacket struct {
	Type   PacketType
	Serial uint16
}

type LoginPacket struct {
	IMEI   string
	Serial uint16
}

type LocationPacket struct {
	Serial uint16
}

func (lp *LoginPacket) ReceiveAck(ack *AckPacket) bool {
	return ack.Type == LoginType && ack.Serial == lp.Serial
}

func (lp *LocationPacket) ReceiveAck(ack *AckPacket) bool {
	return ack.Type == LocationType && ack.Serial == lp.Serial
}
