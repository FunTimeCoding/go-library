package strings

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestElementAfter(t *testing.T) {
	assert.String(
		t,
		"b",
		ElementAfter([]string{"a", "b"}, "a"),
	)
}
