package tcp

func (c *Client) writeLoop() {
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		raw := <-c.sendCh // wait until sendCh

		_, err := c.conn.Write(raw) // configuration
		if err != nil {
			c.logger.Error("Connection returned error", "err", err)
			c.cancel()
			return
		}
	}
}
