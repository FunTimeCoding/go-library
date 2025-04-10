package alert

import "slices"

func (a *Alert) HasTag(s string) bool {
	return slices.Contains(a.Tags, s)
}
