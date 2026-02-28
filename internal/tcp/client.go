package tcp

import (
	"context"
	"log/slog"
	"net"
	"sync"
)

type Client struct {
	conn net.Conn

	sendCh chan []byte // buffered (32) - make(chan []byte, 32)
	AckCh  chan []byte

	readBuf []byte

	ctx    context.Context
	cancel context.CancelFunc

	wg sync.WaitGroup

	logger *slog.Logger

	OnHandleAck func([]byte)
}

func NewClient(conn net.Conn) *Client {
	logger := slog.Default().With(slog.String("layer", "tcp.client"))

	return &Client{
		conn:   conn,
		sendCh: make(chan []byte, 32),
		AckCh:  make(chan []byte),
		logger: logger,
	}
}
