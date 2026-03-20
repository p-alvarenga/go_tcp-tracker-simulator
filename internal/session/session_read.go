package session

import "fmt"

func (s *Session) readLoop() error {
	buf := make([]byte, 4096)

	for {
		select {
		case <-s.ctx.Done():
			s.shutdown()
			return s.ctx.Err()
		default:
		}

		n, err := s.conn.Read(buf)
		if err != nil {
			s.shutdown()
			return err
		}

		data := make([]byte, n)
		copy(data, buf[:n])

		s.handleIncoming(data)
	}
}

func (s *Session) handleIncoming(raw []byte) {
	// 1. Extract frame
	// 2. Decode frame into ACKPacket
	// 3. Calls translateACK()
	// 4. Emit event generated with translateACK()

	s.logger.Info("Received", "raw", fmt.Sprintf("% X", raw))
}
