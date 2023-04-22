package math

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestLogarithm(t *testing.T) {
	assert.Float(t, 0, Logarithm(1))
}
