package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestSwapInList(t *testing.T) {
	elements := []string{Alfa, Bravo, Charlie}
	Swap(elements, 1, 2)
	assert.Any(t, []string{Alfa, Charlie, Bravo}, elements)
}
