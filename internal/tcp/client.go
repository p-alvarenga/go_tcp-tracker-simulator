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
	ReadCh chan []byte

	readBuf []byte

	ctx    context.Context
	cancel context.CancelFunc

	wg sync.WaitGroup

	logger *slog.Logger
}

func NewClient(conn net.Conn, rootLogger *slog.Logger) *Client { // created by Simulator
	return &Client{
		conn:   conn,
		sendCh: make(chan []byte, 32),
		ReadCh: make(chan []byte),
		logger: rootLogger.With(slog.String("layer", "tcp.client")),
	}
}
