package key_reader

import "time"

func (r *Reader) checker() {
	go func() {
		t := time.NewTicker(10 * time.Millisecond)
		defer t.Stop()

		for range t.C {
			r.check()
		}
	}()
}
