package protocol

type Packet interface {
	Build() ([]byte, error)

	Kind() PacketType
	GetSerial() uint16

	IsValid() bool
}
