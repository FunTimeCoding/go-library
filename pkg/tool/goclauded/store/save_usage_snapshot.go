package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/usage_snapshot"
)

func (s *Store) SaveUsageSnapshot(
	sessionPercent int,
	weeklyPercent int,
	sonnetPercent int,
	creditPercent int,
) {
	errors.PanicOnError(
		s.database.Create(
			usage_snapshot.New(
				sessionPercent,
				weeklyPercent,
				sonnetPercent,
				creditPercent,
				s.clock(),
			),
		).Error,
	)
}
