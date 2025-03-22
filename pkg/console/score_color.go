package console

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/math/in_range"
	"github.com/funtimecoding/go-library/pkg/math/range_mapping"
)

func ScoreColor(
	v []*range_mapping.Mapping,
	i face.ScoreColorable,
) {
	for _, m := range v {
		if in_range.LeftOpen(i.Score(), m.Range) {
			switch m.Value {
			case GreenColor:
				i.SetScoreColorFunction(Green)
			case YellowColor:
				i.SetScoreColorFunction(Yellow)
			case RedColor:
				i.SetScoreColorFunction(Red)
			}

			break
		}
	}

	if i.ScoreColorFunction() == nil {
		i.SetScoreColorFunction(Red)
	}
}
