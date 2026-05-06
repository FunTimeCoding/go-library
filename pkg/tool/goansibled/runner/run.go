package runner

import "time"

func (r *Runner) run() {
	r.recovery.Run(r.gitClone)
	r.recovery.Run(r.apply)
	syncTicker := time.NewTicker(SyncInterval)
	defer syncTicker.Stop()
	applyTicker := time.NewTicker(ApplyInterval)
	defer applyTicker.Stop()

	for {
		select {
		case <-r.stop:
			return
		case <-syncTicker.C:
			changed := false
			r.recovery.Run(func() { changed = r.gitSync() })

			if changed {
				r.recovery.Run(r.apply)
				applyTicker.Reset(ApplyInterval)
			}
		case <-applyTicker.C:
			r.recovery.Run(r.apply)
		}
	}
}
