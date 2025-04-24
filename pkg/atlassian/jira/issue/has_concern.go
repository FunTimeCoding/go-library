package issue

import "slices"

func (i *Issue) HasConcern(c string) bool {
	return slices.Contains(i.concerns, c)
}
