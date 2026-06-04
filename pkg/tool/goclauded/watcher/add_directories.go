package watcher

import (
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (w *Watcher) addDirectories(n *fsnotify.Watcher) {
	for _, d := range sourceDirectories() {
		errors.PanicOnError(n.Add(d))
	}
}
