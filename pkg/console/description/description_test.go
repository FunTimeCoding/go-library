package description

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestDescription(t *testing.T) {
	assert.Any(
		t,
		&Description{Title: "Alfa", Short: "Bravo"},
		New(strings.Alfa, strings.Bravo),
	)
}
