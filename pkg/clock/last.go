package clock

import "time"

func (c *Clock) Last() time.Time {
	return c.last
}
