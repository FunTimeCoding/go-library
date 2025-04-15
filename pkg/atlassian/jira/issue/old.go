package issue

import "time"

func (i *Issue) Old() bool {
	if d := time.Since(i.ChangeTime()); d > 4*7*24*time.Hour {
		return true
	}

	return false
}
