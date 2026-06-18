package watcher

import "github.com/funtimecoding/go-library/pkg/errors"

func (w *Watcher) Stop() {
	if w.notifier != nil {
		errors.PanicOnError(w.notifier.Close())
	}
}
