package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestToBoolean(t *testing.T) {
	assert.True(t, ToBoolean(BooleanTrue))
	assert.False(t, ToBoolean(BooleanFalse))
}
