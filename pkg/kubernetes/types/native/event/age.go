package event

import "time"

func (e *Event) Age() time.Duration {
	return time.Since(*e.Create)
}
