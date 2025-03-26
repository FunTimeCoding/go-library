package score_colorer

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/math/in_range"
)

func (c *Colorer) Set(i face.ScoreColorable) {
	value := i.Score() / c.largest

	for _, m := range c.mapping {
		if in_range.LeftOpen(value, m.Range) {
			if f, okay := c.assignments[m.Value]; okay {
				i.SetScoreColor(f)
			}

			break
		}
	}

	if first := c.mapping[0]; value <= first.Range.L {
		i.SetScoreColor(console.Green)
	}

	if last := c.mapping[len(c.mapping)-1]; value >= last.Range.R {
		i.SetScoreColor(console.Red)
	}

	if i.ScoreColor() == nil {
		panic("unable to determine color")
	}
}
