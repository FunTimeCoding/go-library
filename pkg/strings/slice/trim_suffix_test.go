package slice

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"testing"
)

func TestTrimSuffix(t *testing.T) {
	assert.Strings(
		t,
		[]string{"a", "b"},
		TrimSuffix([]string{"a.", "b."}, separator.Dot),
	)
}
