package hypertext

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/fixture"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"testing"
)

func TestDivisions(t *testing.T) {
	assert.Any(
		t,
		[]string{"Example DT", "Example DD"},
		Divisions(
			Document(fixture.File(constant.HypertextPath, "test.html")),
		),
	)
}
