package issue

import "time"

func (i *Issue) NewerThanHours(f int) bool {
	if age := time.Since(i.ChangeTime()).Hours(); int(age) < f {
		return true
	}

	return false
}
