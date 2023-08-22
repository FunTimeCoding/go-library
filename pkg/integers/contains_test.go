package integers

import (
	assert2 "github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestContains(t *testing.T) {
	assert2.True(t, Contains([]int{0}, 0))
	assert2.False(t, Contains([]int{0}, 1))
}
