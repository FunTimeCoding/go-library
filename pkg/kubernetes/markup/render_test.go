package markup

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestRender(t *testing.T) {
	assert.Any(t, "null\n", Render(nil))
}
