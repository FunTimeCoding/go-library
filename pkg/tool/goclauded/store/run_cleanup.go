package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
	"time"
)

func (s *Store) RunCleanup() {
	for {
		time.Sleep(1 * time.Hour)
		cutoff := s.clock().Add(-3 * 24 * time.Hour)
		result := s.database.
			Where("last_seen < ?", cutoff).
			Delete(session.New())
		errors.PanicOnError(result.Error)

		if result.RowsAffected > 0 {
			s.notify()
		}
	}
}
