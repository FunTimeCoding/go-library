package list

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestJoin(t *testing.T) {
	assert.Any(
		t,
		[]string{"Alfa", "Bravo"},
		Join([][]string{{upper.Alfa}, {upper.Bravo}}),
	)
}
