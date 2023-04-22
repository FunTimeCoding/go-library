package object_notation

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestComment(t *testing.T) {
	assert.String(t, "a:\n\"b\"", Comment("a", "b"))
}
