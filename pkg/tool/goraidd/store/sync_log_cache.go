package store

import (
	"github.com/funtimecoding/go-library/pkg/gw2"
	gw2Constant "github.com/funtimecoding/go-library/pkg/gw2/constant"
	"github.com/funtimecoding/go-library/pkg/gw2/log_manager/log"
	"github.com/funtimecoding/go-library/pkg/raid"
	"github.com/funtimecoding/go-library/pkg/system"
	"log/slog"
	"os"
	"path/filepath"
)

func (s *Store) syncLogCache() {
	path := filepath.Join(s.logCachePath, gw2Constant.LogFile)

	if _, e := os.Stat(path); e != nil {
		slog.Warn("log cache not found", "path", path)

		return
	}

	b := system.ReadBytes(s.logCachePath, gw2Constant.LogFile)
	logs := log.NewSlice(gw2.ParseLogs(b, false))
	count := 0

	for _, l := range logs {
		fight := raid.Fight{
			Filename:    l.Raw.FileName,
			Timestamp:   l.Time,
			MapID:       l.Raw.MapID,
			AlliedCount: len(l.Raw.Players),
		}
		result := s.mapper.
			Where("filename = ?", fight.Filename).
			FirstOrCreate(&fight)

		if result.RowsAffected > 0 {
			count++
		}
	}

	slog.Info("log cache synced", "total", len(logs), "new", count)
}
