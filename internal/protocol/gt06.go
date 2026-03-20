package protocol

type PacketType byte

const (
	LoginType    PacketType = 0x01
	LocationType PacketType = 0x12
)

var (
	startBytes = []byte{0x78, 0x78}
	stopBytes  = []byte{0x0D, 0x0A}
)
