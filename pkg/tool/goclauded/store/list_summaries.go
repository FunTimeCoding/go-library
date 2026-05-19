package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/summary"
)

func (s *Store) ListSummaries() []summary.Summary {
	var result []summary.Summary
	errors.PanicOnError(s.database.Find(&result).Error)

	return result
}
