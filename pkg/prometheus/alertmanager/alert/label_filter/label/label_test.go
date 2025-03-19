package label

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestLabel(t *testing.T) {
	assert.True(
		t,
		New(strings.Alfa, strings.Bravo).Match("Alfa", "Bravo"),
	)
}
