package runner

import (
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/store"
	"time"
)

func (r *Runner) run() {
	r.recovery.Run(r.gitClone)
	r.recovery.Run(func() { r.apply(r.playbook, store.TriggerTimer) })
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
				r.recovery.Run(
					func() { r.apply(r.playbook, store.TriggerTimer) },
				)
				applyTicker.Reset(ApplyInterval)
			}
		case <-applyTicker.C:
			r.recovery.Run(
				func() { r.apply(r.playbook, store.TriggerTimer) },
			)
		case request := <-r.trigger:
			playbooks := request.Playbook

			if len(playbooks) == 0 {
				playbooks = r.playbook
			}

			r.recovery.Run(
				func() { r.apply(playbooks, store.TriggerManual) },
			)
		case <-cleanupTicker.C:
			r.recovery.Run(r.store.Cleanup)
		}
	}
}
