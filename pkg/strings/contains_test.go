package strings

import (
	assert2 "github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestContains(t *testing.T) {
	assert2.True(t, Contains([]string{Alfa}, Alfa))
	assert2.False(t, Contains([]string{Alfa}, Bravo))
}
