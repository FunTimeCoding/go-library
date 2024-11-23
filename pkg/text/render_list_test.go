package text

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestRenderList(t *testing.T) {
	assert.String(
		t,
		RenderList("title", []string{"element"}),
		"title:\nelement",
	)
}
