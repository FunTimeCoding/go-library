package service

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/tracker"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"path/filepath"
)

func (s *Service) EnrichSession(identifier string) {
	path := filepath.Join(
		s.harbor,
		join.Empty(identifier, constant.NotationLogExtension),
	)
	s.statesMu.Lock()
	state, exists := s.states[identifier]

	if !exists {
		state = tracker.New()
	}

	s.statesMu.Unlock()

	if e := tracker.Read(path, state); e != nil {
		s.reporter.CaptureException(e)

		return
	}

	s.RefreshSession(identifier, state)
	s.statesMu.Lock()
	s.states[identifier] = state
	s.statesMu.Unlock()
}
