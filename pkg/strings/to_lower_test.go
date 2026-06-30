package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestToLower(t *testing.T) {
	assert.Strings(
		t,
		[]string{"alfa", "bravo"},
		ToLower([]string{upper.Alfa, upper.Bravo}),
	)
}
