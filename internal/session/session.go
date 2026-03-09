package session

import (
	"context"
	"log/slog"
	"net"
	"sync"
	"time"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/protocol/gt06"
)

type Session struct {
	conn        net.Conn
	addr        string
	connTimeout time.Duration

	sendCh chan []byte // buffered (32) - make(chan []byte, 32)
	ACKCh  chan *gt06.ACKPacket
	done   chan struct{}

	readBuf []byte

	ctx    context.Context
	cancel context.CancelFunc

	wg sync.WaitGroup

	logger *slog.Logger
}

func New(addr string, timeout time.Duration, rootLogger *slog.Logger) (*Session, error) {
	return &Session{
		addr:        addr,
		connTimeout: timeout,

		sendCh: make(chan []byte, 32),
		ACKCh:  make(chan *gt06.ACKPacket),
		done:   make(chan struct{}),

		logger: rootLogger.With(slog.String("layer", "Session")),
	}, nil
}

func (s *Session) Start(ctx context.Context) {
	s.ctx, s.cancel = context.WithCancel(ctx)

	err := s.TryConnect()
	if err != nil {
		s.logger.Error("Could not connect", "err", err)
		s.cancel()
		return
	}

	go s.readLoop()
	go s.writeLoop()

	<-s.ctx.Done()

	s.conn.Close()
	close(s.done)
}

func (s *Session) TryConnect() error {
	var err error
	s.conn, err = net.DialTimeout("tcp", s.addr, s.connTimeout)
	return err
}

func (s *Session) Done() <-chan struct{} {
	return s.done
}
