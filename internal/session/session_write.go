package session

import (
	"fmt"
	"go_tcp-tracker-simulator/internal/protocol"
)

func (s *Session) writeLoop() error {
	for {
		select {
		case <-s.ctx.Done():
			s.shutdown()
			return fmt.Errorf("Connection closed")
		case raw := <-s.out:
			s.logger.Info("Sending", "raw", fmt.Sprintf("% 0X", raw))
			s.conn.Write(raw)
		}
	}
}

func (s *Session) sendPacket(pkt protocol.Packet) error {
	if !pkt.IsValid() {
		return fmt.Errorf("Received packet is not valid %s", pkt)
	}

	raw, err := pkt.Build()
	if err != nil {
		return fmt.Errorf("Could not build packet err=%s", err)
	}

	s.out <- raw

	return nil
}
