package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestUnsignedToInteger(t *testing.T) {
	assert.Unsigned(t, 5, ToUnsignedInteger(Five, 0))
	assert.Unsigned(t, 5, ToUnsignedInteger(" 5", 0))
}

func TestUnsignedToIntegerStrict(t *testing.T) {
	assert.Unsigned(t, 1, ToUnsignedIntegerStrict(One))
	assert.Unsigned(t, 1, ToUnsignedIntegerStrict(" 1"))
}
