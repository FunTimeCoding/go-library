package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/completion"
)

func (s *Store) UpdateCompletionSummary(
	sessionIdentifier string,
	topic string,
	message string,
) {
	errors.PanicOnError(
		s.database.Model(completion.Stub()).
			Where(
				"session_identifier = ? AND topic = ?",
				sessionIdentifier,
				topic,
			).
			Update("summary", message).Error,
	)
	s.notify()
}
