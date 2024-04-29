package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestSliceTrimSuffix(t *testing.T) {
	assert.Strings(
		t,
		[]string{"a", "b"},
		SliceTrimSuffix([]string{"a.", "b."}, "."),
	)
}
