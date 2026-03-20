package session

import (
	"context"
	"fmt"
	"go_tcp-tracker-simulator/internal/config"
	"go_tcp-tracker-simulator/internal/domain"
	"go_tcp-tracker-simulator/internal/protocol"
	"log/slog"
	"net"
	"sync"
	"time"
)

type Session struct {
	addr string
	conn net.Conn

	cfg          *config.SessionConfig
	retryCounter uint

	id    domain.SessionID
	imei  domain.IMEI
	state domain.SessionState

	sink       chan<- SessionEvent
	out        chan []byte
	lastPacket protocol.Packet

	closeOnce sync.Once
	ctx       context.Context
	cancel    context.CancelFunc

	logger *slog.Logger
}

func newSession(addr string, id domain.SessionID, imei domain.IMEI, sink chan<- SessionEvent, cfg *config.SessionConfig, ctx context.Context, log *slog.Logger) *Session {
	s := &Session{
		addr: addr,

		cfg:          cfg,
		retryCounter: 0,

		id:    id,
		imei:  imei,
		state: domain.StateSessionCreated,

		sink: sink,
		out:  make(chan []byte, 32),
	}

	s.logger = log.With("lyr", "Session", "id", s.id)
	s.ctx, s.cancel = context.WithCancel(ctx)

	return s
}

func (s *Session) run() {
	s.logger.Info("Started session")

	for {
		if s.cfg.Reconnection.MaxRetries > 0 && s.retryCounter >= s.cfg.Reconnection.MaxRetries {
			s.emitError(fmt.Errorf("Retry counter maxed up"))
			s.shutdown()
			return
		}

		ok := s.connect()
		if !ok {
			s.retryCounter++

			if !s.waitRetry(s.nextRetryDelay()) {
				s.emitError(fmt.Errorf("Context closed"))
			}
			continue
		}

		s.setState(domain.StateSessionConnected)
		s.logger.Info("Connected")
		s.retryCounter = 0
		s.emit(domain.EventSessionConnected)

		s.runIO()
	}
}

func (s *Session) runIO() {
	errIO := make(chan error, 2)

	go func() { errIO <- s.readLoop() }()
	go func() { errIO <- s.writeLoop() }()

	select {
	case <-s.ctx.Done():
		s.shutdown()
		s.emitError(fmt.Errorf("Context end"))

	case err := <-errIO:
		s.shutdown()
		s.emitError(err)
	}
}

func (s *Session) connect() bool {
	var err error

	s.conn, err = net.DialTimeout("tcp", s.addr, s.cfg.DialTimeout)
	if err != nil {
		s.logger.Warn("Could not connect into address", "addr", s.addr)
		return false
	}

	s.closeOnce = sync.Once{}
	return true
}

func (s *Session) shutdown() {
	s.closeOnce.Do(func() {
		s.logger.Warn("Session shutdown")

		s.cancel()
		if s.conn != nil {
			s.conn.Close()
			s.conn = nil
		}
	})
}

func (s *Session) emit(kind domain.SessionEventType) {
	s.sink <- SessionEvent{
		Kind:      kind,
		SessionID: s.id,
		IMEI:      s.imei,
		Error:     nil,
		Time:      time.Now(),
	}
}

func (s *Session) emitError(err error) {
	s.sink <- SessionEvent{
		Kind:      domain.EventSessionError,
		SessionID: s.id,
		Error:     err,
		Time:      time.Now(),
	}
}

func (s *Session) setState(st domain.SessionState) { s.state = st }
