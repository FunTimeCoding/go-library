package service

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/tracker"
	"os"
	"path/filepath"
	"strings"
)

func (s *Service) PopulateCache() {
	entries, e := os.ReadDir(s.harbor)

	if e != nil {
		s.reporter.CaptureException(e)

		return
	}

	s.statesMu.Lock()
	defer s.statesMu.Unlock()

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(
			entry.Name(),
			constant.NotationLogExtension,
		) {
			continue
		}

		identifier := strings.TrimSuffix(
			entry.Name(),
			constant.NotationLogExtension,
		)
		state, exists := s.states[identifier]

		if !exists {
			state = tracker.New()
		}

		path := filepath.Join(s.harbor, entry.Name())

		if tracker.Read(path, state) != nil {
			continue
		}

		s.states[identifier] = state
	}
}
