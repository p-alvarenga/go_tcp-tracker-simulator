package tcp

import "context"

func (c *Client) Start(parentCtx context.Context) {
	c.ctx, c.cancel = context.WithCancel(parentCtx)

	c.wg.Add(2)

	go c.readLoop()
	go c.writeLoop()

	<-c.ctx.Done()

	c.wg.Wait()
	c.conn.Close()
}
