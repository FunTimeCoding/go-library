package slice

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestTrimSuffix(t *testing.T) {
	assert.Strings(
		t,
		[]string{"a", "b"},
		TrimSuffix([]string{"a.", "b."}, "."),
	)
}
