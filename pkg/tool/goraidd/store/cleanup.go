package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

func (s *Store) cleanup() {
	cutoff := time.Now().AddDate(0, 0, -14)
	entries, e := os.ReadDir(s.eliteInsightsPath)

	if e != nil {
		slog.Warn("cleanup: cannot read elite insights dir", "error", e)

		return
	}

	deleted := 0

	for _, entry := range entries {
		info, e := entry.Info()

		if e != nil {
			continue
		}

		if info.ModTime().After(cutoff) {
			continue
		}

		errors.PanicOnError(
			os.Remove(
				filepath.Join(s.eliteInsightsPath, entry.Name()),
			),
		)
		deleted++
	}

	if deleted > 0 {
		slog.Info("cleanup completed", "deleted", deleted)
	}
}
