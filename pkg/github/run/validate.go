package run

import "slices"

func (r *Run) Validate() {
	if len(r.Jobs) > 0 && r.Jobs[0].Fail() {
		if !slices.Contains(r.concern, Failed) {
			r.concern = append(r.concern, Failed)
		}
	}
}
