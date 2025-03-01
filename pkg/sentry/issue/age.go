package issue

import "time"

func (i *Issue) Age() time.Duration {
	return time.Since(i.Create)
}
