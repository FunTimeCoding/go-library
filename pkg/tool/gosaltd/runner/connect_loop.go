package runner

import "time"

func (r *Runner) connectLoop() bool {
	for r.salt == nil {
		r.recovery.Run(r.connect)

		if r.salt != nil {
			return true
		}

		select {
		case <-r.stop:
			return false
		case <-time.After(KeySyncInterval):
		}
	}

	return true
}
