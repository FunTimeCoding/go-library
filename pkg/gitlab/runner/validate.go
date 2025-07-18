package runner

import "slices"

func (r *Runner) Validate() {
	if r.Paused && !slices.Contains(r.concern, Paused) {
		r.concern = append(r.concern, Paused)
	}

	if r.Address == "" && !slices.Contains(r.concern, NoAddressConcern) {
		r.concern = append(r.concern, NoAddressConcern)
	}
}
