package issue

import "slices"

func (i *Issue) IsStatusAny(s ...string) bool {
	return slices.Contains(s, i.Status)
}
