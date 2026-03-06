package session

func (s *Session) writeLoop() {
	for {
		select {
		case <-s.ctx.Done():
			return
		default:
		}

		raw := <-s.sendCh // wait until sendCh

		_, err := s.conn.Write(raw) // configuration
		if err != nil {
			s.logger.Error("Connection returned error", "err", err)
			s.cancel()
			return
		}
	}
}
