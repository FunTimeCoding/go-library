package age_colorer

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/age_colorer/age_fixture"
	"github.com/funtimecoding/go-library/pkg/math/range_mapping"
	"github.com/funtimecoding/go-library/pkg/math/ranges"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"testing"
	"time"
)

func TestColorer(t *testing.T) {
	g := age_fixture.New(0 * time.Hour)
	y := age_fixture.New(15 * time.Hour)
	r := age_fixture.New(30 * time.Hour)

	assert.Float(
		t,
		0,
		timeLibrary.HoursToDecades(g.Age().Hours()),
	)
	assert.Float(
		t,
		0.057534246575342465,
		timeLibrary.HoursToDecades(y.Age().Hours()),
	)
	assert.Float(
		t,
		0.11506849315068493,
		timeLibrary.HoursToDecades(r.Age().Hours()),
	)

	c := Default(g, y, r)
	assert.Any(
		t,
		[]*range_mapping.Mapping{
			{
				Range: ranges.Range{
					L: 0,
					R: 0.038356164383561646,
				},
				Value: "green",
			},
			{
				Range: ranges.Range{
					L: 0.038356164383561646,
					R: 0.07671232876712329,
				},
				Value: "yellow",
			},
			{
				Range: ranges.Range{
					L: 0.07671232876712329,
					R: 0.11506849315068493,
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

	assert.String(t, console.Green("g"), g.AgeColor()("g"))
	assert.String(t, console.Yellow("y"), y.AgeColor()("y"))
	assert.String(t, console.Red("r"), r.AgeColor()("r"))
}
