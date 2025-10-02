package alert

import "slices"

func (a *Alert) Validate() {
	if a.Suppressed() && !slices.Contains(a.concern, Silent) {
		a.concern = append(a.concern, Silent)
	}
}
