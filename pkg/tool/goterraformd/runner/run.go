package runner

import (
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/store"
	"time"
)

func (r *Runner) run() {
	r.recovery.Run(r.gitClone)
	r.recovery.Run(r.terraformInit)
	r.recovery.Run(func() { r.apply(store.TriggerTimer) })
	syncTicker := time.NewTicker(SyncInterval)
	defer syncTicker.Stop()
	applyTicker := time.NewTicker(ApplyInterval)
	defer applyTicker.Stop()
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
				r.recovery.Run(r.terraformInit)
				r.recovery.Run(
					func() { r.apply(store.TriggerTimer) },
				)
				applyTicker.Reset(ApplyInterval)
			}
		case <-applyTicker.C:
			r.recovery.Run(
				func() { r.apply(store.TriggerTimer) },
			)
		case <-r.trigger:
			r.recovery.Run(r.terraformInit)
			r.recovery.Run(
				func() { r.apply(store.TriggerManual) },
			)
		case <-cleanupTicker.C:
			r.recovery.Run(r.store.Cleanup)
		}
	}
}
