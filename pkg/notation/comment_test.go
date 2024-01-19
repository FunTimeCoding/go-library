package notation

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestComment(t *testing.T) {
	assert.String(t, "a:\n\"b\"", Comment("a", "b"))
}
