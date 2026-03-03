package tcp

func (c *Client) writeLoop() {
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		raw := <-c.sendCh // wait until sendCh

		n, err := c.conn.Write(raw) // configuration
		if err != nil {
			c.logger.Error("Connection returned error", "err", err)
		}

		if n != len(raw) {
			c.logger.Warn("Could not send all raw packet")
		}
	}
}
