package manager

import (
	"sync"

	"github.com/harvester/harvester/pkg/controller/master/supportbundle/types"
)

type ManagerStatus struct {
	types.ManagerStatus

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
