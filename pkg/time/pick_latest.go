package time

import "time"

func PickLatest(
	a time.Time,
	b time.Time,
) time.Time {
	if a.After(b) {
		return a
	}

	return b
}
