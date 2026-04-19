package store

import (
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/go-library/pkg/errors"
	gw2Constant "github.com/funtimecoding/go-library/pkg/gw2/constant"
	"log/slog"
	"path/filepath"
	"strings"
)

func (s *Store) RunBackground(stop <-chan struct{}) {
	s.enrichExisting()
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

	for {
		select {
		case <-stop:
			return
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if !event.Has(fsnotify.Create) && !event.Has(fsnotify.Write) {
				continue
			}

			if filepath.Base(event.Name) == gw2Constant.LogFile {
				slog.Info("log cache changed, syncing")
				s.syncLogCache()
			}

			if strings.HasSuffix(event.Name, "_detailed_wvw_kill.json") {
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
