package clock

import "time"

func (c *Clock) Now() time.Time {
	c.last = time.Now()

	return c.last
}
