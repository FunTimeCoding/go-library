package service

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
)

func (s *Service) DeleteLabel(
	sessionIdentifier string,
	name string,
	target string,
	key string,
) string {
	old := s.store.DeleteLabel(sessionIdentifier, key)

	if old == "" {
		return fmt.Sprintf("%s (not set)", key)
	}

	change := fmt.Sprintf("%s %s→ (unset)", key, old)
	s.store.LogEvent(sessionIdentifier, constant.Label, name, target, change)
	s.notify()

	return change
}
