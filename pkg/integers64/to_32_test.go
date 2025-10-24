package integers64

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestTo32(t *testing.T) {
	assert.Integer32(t, 0, To32(0))
}
