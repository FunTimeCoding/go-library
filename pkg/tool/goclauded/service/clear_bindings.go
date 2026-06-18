package service

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
)

func (s *Service) ClearBindings() {
	callsigns := s.store.BoundCallsigns()

	for _, c := range callsigns {
		errors.PanicOnError(
			s.PushQueue(
				c,
				constant.QueueReannounce,
				"Re-announce required: MCP binding lost during service restart. Call announce with your session name and topic to restore.",
			),
		)
	}

	s.store.ClearBindings()
}
