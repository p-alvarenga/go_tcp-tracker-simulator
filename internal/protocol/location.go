package protocol

type LocationPacket struct {
	Serial uint16
}

func (p *LocationPacket) Build() ([]byte, error) { return nil, nil }

func (p *LocationPacket) IsValid() bool     { return true }
func (p *LocationPacket) Kind() PacketType  { return LocationType }
func (p *LocationPacket) GetSerial() uint16 { return p.Serial }
