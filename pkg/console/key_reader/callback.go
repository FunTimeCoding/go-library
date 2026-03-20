package key_reader

import "time"

type Callback struct {
	Press func(
		k rune,
		t time.Time,
	)
	Release func(
		k rune,
		d time.Duration,
	)
}
