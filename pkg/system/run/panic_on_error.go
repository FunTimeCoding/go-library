package run

import "github.com/funtimecoding/go-library/pkg/errors"

func (r *Run) PanicOnError() {
	errors.PanicOnError(r.Error)
}
