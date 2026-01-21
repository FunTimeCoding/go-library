package key_reader

import "time"

func (r *Reader) SetDelay(delay time.Duration) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.releaseDelay = delay
}
