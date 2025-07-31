package item

import "time"

func (i *Item) Age() time.Duration {
	if i.Create == nil {
		return 0
	}

	return time.Since(*i.Create)
}
