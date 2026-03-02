package tcp

func (c *Client) writeLoop() {
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		raw := <-c.sendCh // wait until sendCh

		c.logger.Info("Sending packet", "raw", raw)
		c.conn.Write(raw) // configuration
	}
}
