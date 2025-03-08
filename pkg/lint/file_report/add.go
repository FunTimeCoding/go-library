package file_report

import "github.com/funtimecoding/go-library/pkg/lint/concern"

func (r *Report) Add(c *concern.Concern) {
	r.Concerns = append(r.Concerns, c)
}
