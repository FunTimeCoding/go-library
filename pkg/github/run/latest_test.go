package run

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestLatest(t *testing.T) {
	assert.String(
		t,
		"Charlie",
		Latest(
			[]*Run{
				{
					Name:   strings.Alfa,
					Status: Completed,
					Create: assert.NewDay(0),
				},
				{
					Name:   strings.Bravo,
					Status: Completed,
					Create: assert.NewDay(1),
				},
				{
					Name:   strings.Charlie,
					Status: Completed,
					Create: assert.NewDay(2),
				},
			},
		).Name,
	)
}
