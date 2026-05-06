package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"time"
)

func (s *Store) Cleanup() {
	errors.PanicOnError(
		s.mapper.
			Where("created_at < ?", time.Now().Add(-RetentionAge)).
			Delete(&PlaybookRun{}).Error,
	)
}
