package alert

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestSortByAge(t *testing.T) {
	dayOne := assert.NewDay(1)
	dayTwo := assert.NewDay(2)
	descending := SortByAge(
		[]*Alert{
			{
				Name:  strings.Alfa,
				Start: ptr.To(dayOne),
			},
			{
				Name:  strings.Bravo,
				Start: ptr.To(dayTwo),
			},
		},
		false,
	)
	assert.String(t, "Bravo", descending[0].Name)
	assert.String(t, "Alfa", descending[1].Name)
	ascending := SortByAge(
		[]*Alert{
			{
				Name:  strings.Alfa,
				Start: ptr.To(dayOne),
			},
			{
				Name:  strings.Bravo,
				Start: ptr.To(dayTwo),
			},
		},
		true,
	)
	assert.String(t, "Alfa", ascending[0].Name)
	assert.String(t, "Bravo", ascending[1].Name)
}
