package store

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/usage_snapshot"
	"time"
)

func (s *Store) UsageSnapshots(since time.Duration) []usage_snapshot.Snapshot {
	var result []usage_snapshot.Snapshot
	s.database.
		Where("created_at > ?", s.clock().Add(-since)).
		Order("created_at ASC").
		Find(&result)

	return result
}
