package watcher

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/sweep"

func (w *Watcher) scan() {
	sweep.Run(w.harbor)
	w.service.PopulateCache()
	w.service.BackfillSessions()
	w.service.CheckConsistency()
}
