package manager

import "sync"

type ManagerStatus struct {
	// phase to collect bundle
	Phase string

	// fail to collect bundle
	Error bool

	// error message
	ErrorMessage string

	// progress of the bundle collecting. 0 - 100.
	Progress int

	lock sync.RWMutex
}

func (s *ManagerStatus) SetPhase(phase string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Phase = phase
}

func (s *ManagerStatus) SetError(message string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Error = true
	s.ErrorMessage = message
}

func (s *ManagerStatus) SetProgress(progress int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Progress = progress
}
