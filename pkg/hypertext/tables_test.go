package hypertext

import (
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"testing"
)

func TestTables(t *testing.T) {
	assert.Any(
		t,
		[]string{"Example TD"},
		Tables(
			Document(
				internal.FixtureFile(constant.HypertextPath, "test.html"),
			),
		),
	)
}
