package repository

import "slices"

func (r *Repository) HasConcern(s string) bool {
	return slices.Contains(r.concern, s)
}
