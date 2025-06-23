package hypertext

import (
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"testing"
)

func TestTitle(t *testing.T) {
	assert.String(
		t,
		"Test Title",
		Title(
			Document(
				internal.FixtureFile(constant.HypertextPath, "test.html"),
			),
		),
	)
}
