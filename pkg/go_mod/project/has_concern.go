package project

import "slices"

func (p *Project) HasConcern(s string) bool {
	return slices.Contains(p.concern, s)
}
