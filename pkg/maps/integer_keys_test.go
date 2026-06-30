package maps

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestIntegerKeys(t *testing.T) {
	assert.Integers(
		t,
		[]int{0, 1},
		IntegerKeys(map[int]string{0: upper.Alfa, 1: upper.Bravo}),
	)
}
