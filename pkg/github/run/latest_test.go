package run

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestLatest(t *testing.T) {
	assert.String(
		t,
		"Charlie",
		Latest(
			[]*Run{
				{
					Name:   upper.Alfa,
					Status: Completed,
					Create: assert.NewDay(0),
				},
				{
					Name:   upper.Bravo,
					Status: Completed,
					Create: assert.NewDay(1),
				},
				{
					Name:   upper.Charlie,
					Status: Completed,
					Create: assert.NewDay(2),
				},
			},
		).Name,
	)
}
