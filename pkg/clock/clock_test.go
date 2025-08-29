package clock

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestClock(t *testing.T) {
	assert.NotNil(t, New())
}
