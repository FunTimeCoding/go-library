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

	assert.Float(t, 0, g.Score())
	assert.Float(t, 15, y.Score())
	assert.Float(t, 30, r.Score())

	c := Default(g, y, r)
	assert.Any(
		t,
		[]*range_mapping.Mapping{
			{
				Range: ranges.Range{L: 0, R: 10},
				Value: "green",
			},
			{
				Range: ranges.Range{L: 10, R: 20},
				Value: "yellow",
			},
			{
				Range: ranges.Range{L: 20, R: 30},
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

	assert.String(t, console.Green("g"), g.ScoreColor()("g"))
	assert.String(t, console.Yellow("y"), y.ScoreColor()("y"))
	assert.String(t, console.Red("r"), r.ScoreColor()("r"))
}
