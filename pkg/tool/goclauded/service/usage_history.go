package service

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/usage_snapshot"
	"time"
)

func (s *Service) UsageHistory(
	since time.Duration,
) []usage_snapshot.Snapshot {
	return s.store.UsageSnapshots(since)
}
