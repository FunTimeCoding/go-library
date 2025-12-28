package hypertext

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/fixture"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"testing"
)

func TestTitle(t *testing.T) {
	assert.String(
		t,
		"Test Title",
		Title(
			Document(fixture.File(constant.HypertextPath, "test.html")),
		),
	)
}
