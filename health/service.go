package health

import (
	"sync"

	"github.com/fraym/golog"
)

type componentStatus struct {
	healthy bool
	reason  string
}

type Service struct {
	components map[string]componentStatus

	logger golog.Logger
	mux    sync.RWMutex
}

func New(logger golog.Logger) *Service {
	return &Service{
		components: make(map[string]componentStatus),
		logger:     logger,
	}
}

func (s *Service) isHealthy() (bool, string) {
	for _, component := range s.components {
		if !component.healthy {
			return false, component.reason
		}
	}
	return true, ""
}

func (s *Service) IsHealthy(component ...string) (bool, string) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	if len(component) == 0 {
		return s.isHealthy()
	}

	for _, c := range component {
		if ch, ok := s.components[c]; ok && !ch.healthy {
			return false, s.components[c].reason
		}
	}
	return true, ""
}

func (s *Service) SetHealthy(component string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	componentHealth := componentStatus{
		healthy: true,
	}
	s.components[component] = componentHealth

	s.logger.Warn().WithField("component", component).Write("Component is healthy (again)")
}

func (s *Service) SetUnhealthy(component string, reason string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	componentHealth := componentStatus{
		healthy: false,
		reason:  reason,
	}
	s.components[component] = componentHealth

	s.logger.Warn().WithFields(map[string]any{
		"component": component,
		"reason":    reason,
	}).Write("Component became unhealthy")
}
