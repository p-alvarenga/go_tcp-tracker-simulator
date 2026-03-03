package tcp

import (
	"context"
	"log/slog"
	"net"
	"sync"
	"time"
)

type Client struct {
	conn        net.Conn
	addr        string
	connTimeout time.Duration

	sendCh chan []byte // buffered (32) - make(chan []byte, 32)
	ReadCh chan []byte
	done   chan struct{}

	readBuf []byte

	ctx    context.Context
	cancel context.CancelFunc

	wg sync.WaitGroup

	logger *slog.Logger
}

func NewClient(addr string, timeout time.Duration, rootLogger *slog.Logger) (*Client, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:        conn,
		addr:        addr,
		connTimeout: timeout,

		sendCh: make(chan []byte, 32),
		ReadCh: make(chan []byte),
		done:   make(chan struct{}),

		logger: rootLogger.With(slog.String("layer", "Client")),
	}, nil
}

func (c *Client) Start(parentCtx context.Context) {
	c.ctx, c.cancel = context.WithCancel(parentCtx)

	go c.readLoop()
	go c.writeLoop()

	<-c.ctx.Done() // wait until c.cancel()

	c.conn.Close()
	close(c.done)
}

func (c *Client) TryConnect() error {
	var err error
	c.conn, err = net.DialTimeout("tcp", c.addr, c.connTimeout)
	return err
}

func (c *Client) Done() <-chan struct{} {
	return c.done
}
