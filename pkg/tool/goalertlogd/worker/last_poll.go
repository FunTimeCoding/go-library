package worker

import "time"

func (w *Worker) LastPoll() time.Time {
	v := w.lastPoll.Load()

	if v == nil {
		return time.Time{}
	}

	return v.(time.Time)
}
