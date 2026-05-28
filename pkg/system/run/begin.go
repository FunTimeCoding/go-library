package run

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/system/run/process"
)

func (r *Run) Begin(s ...string) face.Process {
	c := r.build(s...)
	p := process.NewPiped(c)
	errors.PanicOnError(c.Start())

	return p
}
