package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/completion"
)

func (s *Store) ListCompletions() []completion.Completion {
	var result []completion.Completion
	errors.PanicOnError(
		s.database.Order("created_at ASC").Find(&result).Error,
	)

	return result
}
