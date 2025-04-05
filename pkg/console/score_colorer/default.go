package score_colorer

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/color_assignment"
	"github.com/funtimecoding/go-library/pkg/face"
)

func Default[T face.ScoreColorable](v ...T) *Colorer {
	return New(
		[]*color_assignment.Assignment{
			color_assignment.New(console.GreenColor, console.Green),
			color_assignment.New(console.YellowColor, console.Yellow),
			color_assignment.New(console.RedColor, console.Red),
		},
		v...,
	)
}
