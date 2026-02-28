package tcp

import (
	"gt06_sim/internal/protocol/gt06"
	"log/slog"
)

func (c *Client) readLoop() {
	buf := make([]byte, 4*1024)

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		n, err := c.conn.Read(buf)
		if err != nil {
			c.logger.Warn("Could not read", slog.Any("err", err))

			c.cancel()
			return
		}

		c.readBuf = buf[:n]
		c.processBuffer()
	}
}

func (c *Client) processBuffer() {
	var frame []byte
	var ok bool
	for {
		frame, c.readBuf, ok = gt06.ExtractFrame(c.readBuf)

		if !ok {
			return
		}

		c.AckCh <- frame
	}
}
