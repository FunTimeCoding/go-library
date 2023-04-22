package ranges

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestFloat(t *testing.T) {
	assert.Float(t, 3, Float([]float64{1, 2, 3}, 2))
}
