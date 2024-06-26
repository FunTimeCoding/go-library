package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestToInteger(t *testing.T) {
	assert.Integer(t, 5, ToInteger("5", 0))
	assert.Integer(t, 5, ToInteger(" 5", 0))
}

func TestToIntegerStrict(t *testing.T) {
	assert.Integer(t, 1, ToIntegerStrict("1"))
	assert.Integer(t, 1, ToIntegerStrict(" 1"))
}
func TestToInteger64Strict(t *testing.T) {
	assert.Integer64(t, 1, ToInteger64Strict("1"))
	assert.Integer64(t, 1, ToInteger64Strict(" 1"))
}
