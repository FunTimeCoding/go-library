package score_colorer

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/color_assignment"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/integers"
)

func Dynamic(
	startHex string,
	endHex string,
	v ...face.ScoreColorable,
) *Colorer {
	count := len(v)
	c := console.Gradient(startHex, endHex, count)
	assignments := make([]*color_assignment.Assignment, count)

	for i := 0; i < count; i++ {
		assignments[i] = color_assignment.New(integers.ToString(i), c[i])
	}

	return New(assignments, v...)
}
