package key_reader

import "time"

func (r *Reader) press(key rune) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	h, okay := r.handlers[key]

	if !okay {
		return
	}

	s := r.states[key]
	t := time.Now()

	if !s.holding {
		s.press = t
		s.holding = true

		if h.Press != nil {
			h.Press(key, t)
		}
	}

	s.lastKey = t
}
