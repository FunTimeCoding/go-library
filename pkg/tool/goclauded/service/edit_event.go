package service

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event"
)

func (s *Service) EditEvent(
	identifier uint,
	message string,
) (*event.Event, error) {
	edited := s.store.EditEvent(identifier, message)

	switch edited.Kind {
	case constant.Summarize:
		s.store.UpdateSummaryBody(edited.SessionIdentifier, message)

		if e := s.pushSummary(edited.Name, message); e != nil {
			return edited, e
		}
	case constant.Complete:
		s.store.UpdateCompletionSummary(
			edited.SessionIdentifier,
			edited.Target,
			message,
		)
	}

	s.notify()

	return edited, nil
}
