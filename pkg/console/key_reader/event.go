package key_reader

import "github.com/nsf/termbox-go"

func (r *Reader) event(e termbox.Event) bool {
	if e.Type != termbox.EventKey {
		return true
	}

	if e.Key == termbox.KeyEsc {
		return false
	}

	k := e.Ch

	if k == 0 {
		k = rune(e.Key)
	}

	r.press(k)

	return true
}
