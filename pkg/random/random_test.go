package random

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestRandom(t *testing.T) {
	assert.True(t, New() != nil)
}
