package strings

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestToBoolean(t *testing.T) {
	assert.Boolean(t, true, ToBoolean("1"))
	assert.Boolean(t, false, ToBoolean("0"))
}
