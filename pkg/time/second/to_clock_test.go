package second

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestToClock(t *testing.T) {
	assert.String(t, "00:00", ToClock(0))
}
