package serverpool

import (
	"sync"

	"github.com/Ciggzy1312/go-lb/backend"
)

type leastConnectionServerPool struct {
	backends []backend.Backend
	mux      sync.RWMutex
}

func (s *leastConnectionServerPool) GetNextValidPeer() backend.Backend {
	s.mux.Lock()
	defer s.mux.Unlock()
	var leastConnBackend backend.Backend
	for _, b := range s.backends {
		if b.IsAlive() {
			leastConnBackend = b
			break
		}
	}

	for _, b := range s.backends {
		if b.IsAlive() && b.GetConnections() < leastConnBackend.GetConnections() {
			leastConnBackend = b
		}
	}
	return leastConnBackend
}

func (s *leastConnectionServerPool) AddBackend(b backend.Backend) {
	s.backends = append(s.backends, b)
}

func (s *leastConnectionServerPool) GetBackends() []backend.Backend {
	return s.backends
}

func (s *leastConnectionServerPool) GetServerPoolSize() int {
	return len(s.backends)
}
