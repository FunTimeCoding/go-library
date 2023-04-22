package interval_crossed

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestFloat(t *testing.T) {
	assert.Boolean(t, false, Float(10, 0, 1))
	assert.Boolean(t, false, Float(10, 1, 9))
	assert.Boolean(t, true, Float(10, 9, 10))
	assert.Boolean(t, true, Float(10, 9, 11))
	assert.Boolean(t, false, Float(10, 10, 11))
}
