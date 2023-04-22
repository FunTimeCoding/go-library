package time

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestSecondsToClock(t *testing.T) {
	assert.String(t, "00:00", SecondsToClock(0))
}
