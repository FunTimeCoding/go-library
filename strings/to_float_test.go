package strings

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestToFloat(t *testing.T) {
	assert.Float(t, 0.5, ToFloat("0.5", 0))
	assert.Float(t, 0.5, ToFloat(" 0.5", 0))
}

func TestToFloatStrict(t *testing.T) {
	assert.Float(t, 1, ToFloatStrict("1.0"))
	assert.Float(t, 1, ToFloatStrict(" 1.0"))
}
