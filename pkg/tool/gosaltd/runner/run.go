package runner

import (
	"github.com/funtimecoding/go-library/pkg/tool/gosaltd/store"
	"time"
)

func (r *Runner) run() {
	r.recovery.Run(r.gitClone)

	if !r.connectLoop() {
		return
	}

	r.recovery.Run(func() { r.apply(store.TriggerTimer) })
	syncTicker := time.NewTicker(SyncInterval)
	defer syncTicker.Stop()
	highstateTicker := time.NewTicker(HighstateInterval)
	defer highstateTicker.Stop()
	keySyncTicker := time.NewTicker(KeySyncInterval)
	defer keySyncTicker.Stop()
	cleanupTicker := time.NewTicker(24 * time.Hour)
	defer cleanupTicker.Stop()

	for {
		select {
		case <-r.stop:
			return
		case <-syncTicker.C:
			changed := false
			r.recovery.Run(func() { changed = r.gitSync() })

			if changed {
				r.recovery.Run(
					func() { r.apply(store.TriggerTimer) },
				)
				highstateTicker.Reset(HighstateInterval)
			}
		case <-highstateTicker.C:
			r.recovery.Run(
				func() { r.apply(store.TriggerTimer) },
			)
		case <-keySyncTicker.C:
			r.recovery.Run(r.keySync)
		case <-r.trigger:
			r.recovery.Run(
				func() { r.apply(store.TriggerManual) },
			)
		case <-cleanupTicker.C:
			r.recovery.Run(r.store.Cleanup)
		}
	}
}
