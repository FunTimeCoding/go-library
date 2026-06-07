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
) (string, error) {
	old, e := s.store.DeleteLabel(sessionIdentifier, key)

	if e != nil {
		return "", e
	}

	if old == "" {
		return fmt.Sprintf("%s (not set)", key), nil
	}

	change := fmt.Sprintf("%s %s→ (unset)", key, old)

	if e := s.store.LogEvent(
		sessionIdentifier,
		constant.Label,
		name,
		target,
		change,
	); e != nil {
		return "", e
	}

	s.notify()

	return change, nil
}
