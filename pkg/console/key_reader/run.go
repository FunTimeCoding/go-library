package key_reader

import "github.com/nsf/termbox-go"

func (r *Reader) Run() {
	r.checker()

	for {
		if e := termbox.PollEvent(); !r.event(e) {
			return
		}
	}
}
