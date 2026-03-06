package session

import "github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol/gt06"

func (s *Session) SendPacket(pkt gt06.Packet) error {
	// validations

	raw, err := pkt.Build()
	if err != nil {
		s.logger.Error("Could not build packet")
	}

	s.sendCh <- raw
	return nil
}
