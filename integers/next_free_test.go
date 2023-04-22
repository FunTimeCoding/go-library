package integers

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestNextFreeNumber(t *testing.T) {
	assert.Integer(t, 0, NextFree([]int{}))
	assert.Integer(t, 1, NextFree([]int{0}))
	assert.Integer(t, 1, NextFree([]int{0, 2}))
}
