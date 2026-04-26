package store

import (
	"github.com/funtimecoding/go-library/pkg/gw2/constant"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (s *Store) enrichExisting() {
	entries, e := os.ReadDir(s.eliteInsightsPath)

	if e != nil {
		slog.Warn("cannot read elite insights dir", "error", e)

		return
	}

	cutoff := time.Now().AddDate(0, 0, -enrichCutoffDays)

	for _, entry := range entries {
		name := entry.Name()

		if !strings.HasSuffix(name, constant.DetailedWvWKillSuffix) {
			continue
		}

		if len(name) < fileDatePrefixLength {
			continue
		}

		fileDate, parseError := time.Parse(
			fileDateLayout,
			name[:fileDatePrefixLength],
		)

		if parseError != nil {
			continue
		}

		if fileDate.Before(cutoff) {
			continue
		}

		s.enrichFile(filepath.Join(s.eliteInsightsPath, name))
	}
}
