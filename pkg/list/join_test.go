package list

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestJoin(t *testing.T) {
	assert.Any(
		t,
		[]string{"Alfa", "Bravo"},
		Join([][]string{{strings.Alfa}, {strings.Bravo}}),
	)
}
