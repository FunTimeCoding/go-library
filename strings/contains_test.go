package strings

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestContains(t *testing.T) {
	assert.True(t, Contains([]string{Alfa}, Alfa))
	assert.False(t, Contains([]string{Alfa}, Bravo))
}
