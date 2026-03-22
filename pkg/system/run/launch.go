package run

import "github.com/funtimecoding/go-library/pkg/errors"

func (r *Run) Launch(s ...string) {
	c := r.build(s...)
	e := c.Start()

	if r.Panic {
		errors.PanicOnError(e)
	}
}
