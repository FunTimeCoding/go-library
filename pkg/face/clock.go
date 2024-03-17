package face

import "time"

type Clock interface {
	IsClock()

	Now() time.Time

	Last() time.Time
}
