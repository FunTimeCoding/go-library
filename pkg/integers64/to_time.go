package integers64

import "time"

func ToTime(i int64) time.Time {
	return time.Unix(i, 0).UTC()
}
