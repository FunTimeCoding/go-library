package score_colorer

import (
	"github.com/funtimecoding/go-library/pkg/console/color_assignment"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/math/range_mapping"
)

func New[T face.ScoreColorable](
	a []*color_assignment.Assignment,
	c ...T,
) *Colorer {
	var largest float64

	for _, e := range c {
		if value := e.Score(); value > largest {
			largest = value
		}
	}

	var values []float64

	for _, e := range c {
		values = append(values, e.Score()/largest)
	}

	var names []string
	assignments := make(map[string]face.SprintFunction)

	for _, e := range a {
		names = append(names, e.Name)
		assignments[e.Name] = e.Function
	}

	return &Colorer{
		assignments: assignments,
		mapping:     range_mapping.Generate(values, names),
		largest:     largest,
	}
}
