package hypertext

import (
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"testing"
)

func TestDivisions(t *testing.T) {
	assert.Any(
		t,
		[]string{"Example DT", "Example DD"},
		Divisions(
			Document(
				internal.FixtureFile(constant.HypertextPath, "test.html"),
			),
		),
	)
}
