package session

import (
	"log/slog"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol/gt06"
)

func (c *Session) readLoop() {
	buf := make([]byte, 4096)

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		n, err := c.conn.Read(buf)
		if err != nil {
			c.logger.Error("Could not read", slog.Any("err", err))
			c.cancel()
			return
		}

		c.readBuf = buf[:n]
		c.frameBuffer()
	}
}

func (c *Session) frameBuffer() {
	var frame []byte
	var ok bool
	for {
		frame, c.readBuf, ok = gt06.ExtractFrame(c.readBuf)
		if !ok {
			return
		}

		c.ReadCh <- frame
	}
}
