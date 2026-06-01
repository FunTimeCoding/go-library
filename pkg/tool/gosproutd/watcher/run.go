package watcher

import (
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
	"path/filepath"
	"time"
)

func (w *Watcher) run() {
	w.recovery.Run(w.scan)
	n, e := fsnotify.NewWatcher()
	errors.PanicOnError(e)
	defer func() { errors.PanicOnError(n.Close()) }()
	w.addDirectories(n)
	sweep := time.NewTicker(10 * time.Minute)
	defer sweep.Stop()
	var lastScan time.Time

	for {
		select {
		case <-w.stop:
			return
		case <-sweep.C:
			w.recovery.Run(func() {
				w.addDirectories(n)
				w.scan()
			})
		case v, okay := <-n.Events:
			if !okay {
				return
			}

			if !v.Has(fsnotify.Create) && !v.Has(fsnotify.Write) &&
				!v.Has(fsnotify.Remove) && !v.Has(fsnotify.Rename) {
				continue
			}

			if time.Since(lastScan) < 1*time.Second {
				continue
			}

			lastScan = time.Now()
			w.recovery.Run(func() {
				w.addDirectories(n)
				w.scan()
			})
		case _, okay := <-n.Errors:
			if !okay {
				return
			}
		}
	}
}

func (w *Watcher) addDirectories(n *fsnotify.Watcher) {
	filepath.Walk(w.seedDirectory, func(
		path string,
		i os.FileInfo,
		e error,
	) error {
		if e != nil {
			return nil
		}

		if i.IsDir() {
			n.Add(path)
		}

		return nil
	})
}
