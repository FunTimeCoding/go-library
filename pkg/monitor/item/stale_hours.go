package item

import "time"

func (i *Item) staleHours() float64 {
	if i.Update == nil {
		return i.ageHours()
	}

	return time.Since(*i.Update).Hours()
}
