package store

import (
	"github.com/funtimecoding/go-library/pkg/gw2"
	"github.com/funtimecoding/go-library/pkg/gw2/constant"
	"github.com/funtimecoding/go-library/pkg/gw2/log_manager/log"
	"github.com/funtimecoding/go-library/pkg/raid"
	"github.com/funtimecoding/go-library/pkg/system"
	"path/filepath"
)

func (s *Store) syncLogCache() {
	system.Stat(filepath.Join(s.logCachePath, constant.LogFile))
	logs := log.NewSlice(
		gw2.ParseLogs(
			system.ReadBytes(s.logCachePath, constant.LogFile),
			false,
		),
	)
	count := 0

	for _, l := range logs {
		f := raid.NewFight()
		f.Filename = l.Raw.FileName
		f.Timestamp = l.Time
		f.MapID = l.Raw.MapID
		f.AlliedCount = len(l.Raw.Players)

		if r := s.mapper.Where(
			"filename = ?",
			f.Filename,
		).FirstOrCreate(f); r.RowsAffected > 0 {
			count++
		}
	}

	s.logger.Structured(
		"log_cache_synced",
		"total",
		len(logs),
		"new",
		count,
	)
}
