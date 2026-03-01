package tcp

import "github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol/gt06"

func (c *Client) SendPacket(pkt gt06.Packet) error {
	// validations
	raw, err := pkt.Build()
	if err != nil {
		return err
	}

	c.sendCh <- raw
	return nil
}
