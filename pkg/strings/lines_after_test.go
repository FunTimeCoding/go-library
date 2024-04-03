package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestLinesAfter(t *testing.T) {
	assert.Strings(
		t,
		[]string{"c"},
		LinesAfter([]string{"a", "b", "c"}, "b"),
	)
}
