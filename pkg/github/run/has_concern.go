package run

import "slices"

func (r *Run) HasConcern(s string) bool {
	return slices.Contains(r.concern, s)
}
