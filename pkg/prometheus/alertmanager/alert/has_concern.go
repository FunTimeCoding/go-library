package alert

import "slices"

func (a *Alert) HasConcern(s string) bool {
	return slices.Contains(a.concern, s)
}
