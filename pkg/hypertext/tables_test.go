package hypertext

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/assert/fixture"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"testing"
)

func TestTables(t *testing.T) {
	assert.Any(
		t,
		[]string{"Example TD"},
		Tables(
			Document(fixture.File(constant.HypertextPath, "test.html")),
		),
	)
}
