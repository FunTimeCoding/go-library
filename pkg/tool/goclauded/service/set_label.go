package service

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
)

func (s *Service) SetLabel(
	sessionIdentifier string,
	name string,
	target string,
	key string,
	value string,
) string {
	old := s.store.SetLabel(sessionIdentifier, key, value)
	var change string

	switch old {
	case value:
		return fmt.Sprintf("%s %s (unchanged)", key, value)
	case "":
		change = fmt.Sprintf("%s (unset)→%s", key, value)
	default:
		change = fmt.Sprintf("%s %s→%s", key, old, value)
	}

	s.store.LogEvent(sessionIdentifier, constant.Label, name, target, change)
	s.notify()

	return change
}
