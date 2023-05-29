package integers

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestContains(t *testing.T) {
	assert.True(t, Contains([]int{0}, 0))
	assert.False(t, Contains([]int{0}, 1))
}
