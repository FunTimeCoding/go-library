package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/completion"
	"time"
)

func (s *Store) RecentCompletions() []completion.Completion {
	var result []completion.Completion
	errors.PanicOnError(
		s.database.
			Where("created_at > ?", s.clock().Add(-24*time.Hour)).
			Order("created_at DESC").
			Find(&result).Error,
	)

	return result
}
