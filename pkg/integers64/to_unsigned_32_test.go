package integers64

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestToUnsigned32(t *testing.T) {
	assert.Unsigned32(t, 0, ToUnsigned32(0))
}
