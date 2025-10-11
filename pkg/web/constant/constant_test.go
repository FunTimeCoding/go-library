package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "User-Agent", UserAgent)
	assert.String(t, "text/markdown", Markdown)
	assert.String(t, "image/x-icon", Icon)
}
