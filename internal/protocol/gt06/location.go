package gt06

import "fmt"

type LocationPacket struct {
	serial uint16
}

func NewLocationPacket() *LocationPacket {
	return &LocationPacket{}
}

func (p *LocationPacket) Build() ([]byte, error) {
	return nil, fmt.Errorf("Not supported yet")
}

func (p *LocationPacket) Type() PacketType {
	return LocationType
}

func (p *LocationPacket) Serial() uint16 {
	return p.serial
}
