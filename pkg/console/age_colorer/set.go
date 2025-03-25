package age_colorer

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/math/in_range"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
)

func (c *Colorer) Set(i face.AgeColorable) {
	decades := timeLibrary.HoursToDecades(i.Age().Hours())

	for _, m := range c.mapping {
		if in_range.LeftOpen(decades, m.Range) {
			if f, okay := c.assignments[m.Value]; okay {
				i.SetAgeColor(f)
			}

			break
		}
	}

	if first := c.mapping[0]; decades <= first.Range.L {
		i.SetAgeColor(console.Green)
	}

	if last := c.mapping[len(c.mapping)-1]; decades >= last.Range.R {
		i.SetAgeColor(console.Red)
	}

	if i.AgeColor() == nil {
		panic("unable to determine color")
	}
}
