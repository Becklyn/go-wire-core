package readyness

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type Service struct {
	components map[string]bool

	logger *logrus.Logger
	mux    sync.RWMutex
}

func New(logger *logrus.Logger) *Service {
	return &Service{
		components: make(map[string]bool),
		logger:     logger,
	}
}

func (s *Service) isReady() (bool, string) {
	for component, isReady := range s.components {
		if !isReady {
			return false, component
		}
	}
	return true, ""
}

func (s *Service) IsReady(component ...string) (bool, string) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	if len(component) == 0 {
		return s.isReady()
	}

	for _, c := range component {
		if ready, ok := s.components[c]; !ok || !ready {
			return false, c
		}
	}
	return true, ""
}

func (s *Service) SetReady(component string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.components[component] = true

	s.logger.WithFields(logrus.Fields{
		"component": component,
	}).Info("Component is ready")
}

func (s *Service) Register(component string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.components[component] = false

	s.logger.WithFields(logrus.Fields{
		"component": component,
	}).Info("Registered unready component")
}
