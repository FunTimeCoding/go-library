package output

import (
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	"strings"
)

type Results struct {
	workDirectory string
	Entries       []*concern.Concern
}

func (r *Results) AddConcern(c *concern.Concern) {
	c.Path = r.Relativize(c.Path)
	r.Entries = append(r.Entries, c)
}

func (r *Results) Relativize(path string) string {
	return strings.TrimPrefix(path, r.workDirectory)
}
