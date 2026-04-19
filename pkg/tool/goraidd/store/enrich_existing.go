package store

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

func (s *Store) enrichExisting() {
	entries, e := os.ReadDir(s.eliteInsightsPath)

	if e != nil {
		slog.Warn("cannot read elite insights dir", "error", e)

		return
	}

	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), "_detailed_wvw_kill.json") {
			s.enrichFile(filepath.Join(s.eliteInsightsPath, entry.Name()))
		}
	}
}
