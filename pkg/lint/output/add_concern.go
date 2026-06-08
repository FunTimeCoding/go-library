package output

import "github.com/funtimecoding/go-library/pkg/lint/concern"

func (r *Results) AddConcern(c *concern.Concern) {
	c.Path = r.Relativize(c.Path)
	r.Entries = append(r.Entries, c)
}
