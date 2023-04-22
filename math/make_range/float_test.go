package make_range

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestFloat(t *testing.T) {
	assert.Any(t, []float64{0, 1}, Float(0, 1))
}
