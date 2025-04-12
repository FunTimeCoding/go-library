package alert

import "time"

func (a *Alert) Age() time.Duration {
	return time.Since(a.Create)
}
