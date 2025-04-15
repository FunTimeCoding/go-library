package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestSwapInList(t *testing.T) {
	actual := []string{Alfa, Bravo, Charlie}
	Swap(actual, 1, 2)
	assert.Any(t, []string{Alfa, Charlie, Bravo}, actual)
}
