package key_reader

import "time"

func New() *Reader {
	return &Reader{
		handlers:     make(map[rune]handler),
		states:       make(map[rune]*state),
		releaseDelay: 200 * time.Millisecond,
	}
}
