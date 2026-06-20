package runner

import (
	"fmt"
	"time"
)

func (r *Runner) run() {
	defer r.drainChannels()

	if r.setupFunction != nil && !r.setupFunction() {
		return
	}

	r.recovery.Run(r.gitClone)

	if r.initFunction != nil {
		r.recovery.Run(r.initFunction)
	}

	r.recovery.Run(func() { r.applyFunction(nil, TriggerTimer) })
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
				if r.initFunction != nil {
					r.recovery.Run(r.initFunction)
				}

				r.recovery.Run(
					func() { r.applyFunction(nil, TriggerTimer) },
				)
				applyTicker.Reset(ApplyInterval)
			}
		case <-applyTicker.C:
			if r.initFunction != nil {
				r.recovery.Run(r.initFunction)
			}

			r.recovery.Run(
				func() { r.applyFunction(nil, TriggerTimer) },
			)
		case request := <-r.sync:
			var result *SyncResult
			r.recovery.Run(func() { result = r.syncWithDiff() })

			if result == nil {
				result = &SyncResult{
					Error: fmt.Errorf("sync failed"),
				}
			}

			request.Response <- result
		case request := <-r.trigger:
			if request.Update {
				r.recovery.Run(func() { r.gitSync() })
			}

			if r.initFunction != nil {
				r.recovery.Run(r.initFunction)
			}

			var value any
			r.recovery.Run(
				func() {
					value = r.applyFunction(
						request.Parameters,
						TriggerManual,
					)
				},
			)

			if request.Response != nil {
				result := &TriggerResult{Value: value}

				if value == nil {
					result.Error = fmt.Errorf("apply failed")
				}

				request.Response <- result
			}
		case <-cleanupTicker.C:
			if r.cleanupFunction != nil {
				r.recovery.Run(r.cleanupFunction)
			}
		}
	}
}
