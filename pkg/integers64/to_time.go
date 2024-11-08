package integers64

import "time"

func ToTime(i int64) time.Time {
	return time.Unix(i/1000, 0).UTC()
}
