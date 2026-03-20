package session

import (
	"context"
	"fmt"
	"go_tcp-tracker-simulator/internal/config"
	"go_tcp-tracker-simulator/internal/domain"
	"go_tcp-tracker-simulator/internal/protocol"
	"log/slog"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
)

type SessionManager struct {
	cfg  *config.SessionConfig
	addr string

	sessions  map[domain.SessionID]*Session
	idCounter atomic.Uint64

	sink chan<- SessionEvent

	mu sync.RWMutex

	ctx    context.Context
	cancel context.CancelFunc

	rootLogger *slog.Logger
}

func NewSessionManager(cfg *config.SessionConfig, sink chan<- SessionEvent, ctx context.Context, logger *slog.Logger) *SessionManager {
	ctx, cancel := context.WithCancel(ctx)

	return &SessionManager{
		cfg:  cfg,
		addr: net.JoinHostPort(cfg.ServerHost, strconv.Itoa(cfg.ServerPort)),

		sessions: make(map[domain.SessionID]*Session),
		sink:     sink,

		ctx:    ctx,
		cancel: cancel,

		rootLogger: logger,
	}
}

func (m *SessionManager) Get(id domain.SessionID) *Session {
	m.mu.RLock()
	s, ok := m.sessions[id]
	m.mu.RUnlock()

	if ok {
		return s
	}

	return nil
}

func (m *SessionManager) CreateAndRun(imei domain.IMEI) (domain.SessionID, error) {
	id := domain.SessionID(m.idCounter.Add(1))

	s := newSession(m.addr, id, imei, m.sink, m.cfg, m.ctx, m.rootLogger)

	m.mu.Lock()
	m.sessions[id] = s
	m.mu.Unlock()

	go s.run()
	return id, nil
}

func (m *SessionManager) SendPacketIntoSession(id domain.SessionID, pkt protocol.Packet) error {
	s := m.Get(id)

	if s == nil {
		return fmt.Errorf("Could not find session %d", id)
	}

	return s.sendPacket(pkt)
}
