package console

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/math/in_range"
	"github.com/funtimecoding/go-library/pkg/math/range_mapping"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
)

func AgeToColor(
	v []*range_mapping.Mapping,
	i face.AgeColorable,
) {
	for _, m := range v {
		if in_range.LeftOpen(
			timeLibrary.HoursToDecades(i.Age().Hours()),
			m.Range,
		) {
			switch m.Value {
			case GreenColor:
				i.SetAgeColorFunction(Green)
			case YellowColor:
				i.SetAgeColorFunction(Yellow)
			case RedColor:
				i.SetAgeColorFunction(Red)
			}

			break
		}
	}

	if i.AgeColorFunction() == nil {
		i.SetAgeColorFunction(Red)
	}
}
