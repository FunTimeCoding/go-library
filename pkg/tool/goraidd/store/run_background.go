package store

import (
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gw2/constant"
	"log/slog"
	"path/filepath"
	"strings"
	"time"
)

func (s *Store) RunBackground(stop <-chan struct{}) {
	s.enrichExisting()
	s.cleanup()
	watcher, e := fsnotify.NewWatcher()
	errors.PanicOnError(e)
	defer func() { errors.PanicOnError(watcher.Close()) }()
	errors.PanicOnError(watcher.Add(s.logCachePath))
	errors.PanicOnError(watcher.Add(s.eliteInsightsPath))
	slog.Info(
		"watching for changes",
		"logCache",
		s.logCachePath,
		"eliteInsights",
		s.eliteInsightsPath,
	)
	cleanupTicker := time.NewTicker(24 * time.Hour)
	defer cleanupTicker.Stop()

	for {
		select {
		case <-stop:
			return
		case <-cleanupTicker.C:
			s.cleanup()
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if !event.Has(fsnotify.Create) && !event.Has(fsnotify.Write) {
				continue
			}

			if filepath.Base(event.Name) == constant.LogFile {
				slog.Info("log cache changed, syncing")
				s.syncLogCache()
			}

			if strings.HasSuffix(event.Name, constant.DetailedWvWKillSuffix) {
				s.enrichFile(event.Name)
			}
		case watchError, ok := <-watcher.Errors:
			if !ok {
				return
			}

			slog.Error("watcher error", "error", watchError)
		}
	}
}
