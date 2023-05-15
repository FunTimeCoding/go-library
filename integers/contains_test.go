package integers

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestContains(t *testing.T) {
	assert.Boolean(t, true, Contains([]int{0}, 0))
	assert.Boolean(t, false, Contains([]int{0}, 1))
}
