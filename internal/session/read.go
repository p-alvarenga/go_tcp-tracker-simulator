package session

import (
	"log/slog"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol/gt06"
)

func (s *Session) readLoop() {
	buf := make([]byte, 4096)

	for {
		select {
		case <-s.ctx.Done():
			return
		default:
		}

		n, err := s.conn.Read(buf)
		if err != nil {
			s.logger.Error("Could not read", slog.Any("err", err))
			s.cancel()
			return
		}

		s.readBuf = buf[:n]
		s.processBuffer()
	}
}

func (s *Session) processBuffer() {
	var frame []byte
	var ok bool
	for {
		frame, s.readBuf, ok = gt06.ExtractFrame(s.readBuf)
		if !ok {
			return
		}

		ack, err := gt06.DecodeACK(frame)
		if err != nil {
			s.logger.Error("Error decoding ACK Packet", "err", err)
			continue
		}

		s.ACKCh <- ack
	}
}
