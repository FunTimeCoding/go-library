package runner

import "slices"

func (r *Runner) HasConcern(s string) bool {
	return slices.Contains(r.concern, s)
}
