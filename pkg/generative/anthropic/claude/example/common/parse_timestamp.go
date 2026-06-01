package common

import "time"

func ParseTimestamp(s string) time.Time {
	t, e := time.Parse(time.RFC3339Nano, s)

	if e == nil {
		return t
	}

	t, e = time.Parse("2006-01-02T15:04:05.000Z", s)

	if e == nil {
		return t
	}

	return time.Time{}
}
