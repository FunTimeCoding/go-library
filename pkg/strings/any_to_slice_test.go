package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestAnyToSlice(t *testing.T) {
	assert.Strings(t, []string{"Alfa"}, AnyToSlice(Alfa))
	assert.Strings(
		t, []string{"Alfa", "Bravo"},
		AnyToSlice([]string{Alfa, Bravo}),
	)
}
