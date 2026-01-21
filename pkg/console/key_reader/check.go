package key_reader

import "time"

func (r *Reader) check() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	t := time.Now()

	for k, s := range r.states {
		if s.holding && t.Sub(s.lastKey) > r.releaseDelay {
			d := t.Sub(s.press)
			s.holding = false

			if h, okay := r.handlers[k]; okay && h.Release != nil {
				h.Release(k, d)
			}
		}
	}
}
