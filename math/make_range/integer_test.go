package make_range

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestInteger(t *testing.T) {
	assert.Any(t, []int{0, 1}, Integer(0, 1))
}
