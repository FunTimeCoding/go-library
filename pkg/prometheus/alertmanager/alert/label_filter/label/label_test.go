package label

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestLabel(t *testing.T) {
	assert.True(
		t,
		New(upper.Alfa, upper.Bravo).Match(upper.Alfa, upper.Bravo),
	)
}
