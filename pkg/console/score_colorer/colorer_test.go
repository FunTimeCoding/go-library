package score_colorer

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/score_colorer/score_fixture"
	"github.com/funtimecoding/go-library/pkg/math/range_mapping"
	"github.com/funtimecoding/go-library/pkg/math/ranges"
	"testing"
)

func TestColorer(t *testing.T) {
	g := score_fixture.New(0)
	y := score_fixture.New(15)
	r := score_fixture.New(30)

	c := Default(g, y, r)
	assert.Any(
		t,
		[]*range_mapping.Mapping{
			{
				Range: ranges.Range{
					L: 0,
					R: 0.3333333333333333,
				},
				Value: "green",
			},
			{
				Range: ranges.Range{
					L: 0.3333333333333333,
					R: 0.6666666666666666,
				},
				Value: "yellow",
			},
			{
				Range: ranges.Range{
					L: 0.6666666666666666,
					R: 1,
				},
				Value: "red",
			},
		},
		c.mapping,
	)

	c.Set(g)
	c.Set(y)
	c.Set(r)

	console.GreenInstance.EnableColor()
	console.YellowInstance.EnableColor()
	console.RedInstance.EnableColor()

	// Not sure if function pointers can be compared, so compare output

	assert.String(
		t,
		console.Green("%s", "g"),
		g.ScoreColor()("g"),
	)
	assert.String(
		t,
		console.Yellow("%s", "y"),
		y.ScoreColor()("y"),
	)
	assert.String(
		t,
		console.Red("%s", "r"),
		r.ScoreColor()("r"),
	)
}
