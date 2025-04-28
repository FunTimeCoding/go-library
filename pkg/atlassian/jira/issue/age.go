package issue

import "time"

func (i *Issue) Age() time.Duration {
	return time.Since(time.Time(i.Raw.Fields.Created))
}
