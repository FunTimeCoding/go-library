package ranges

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestRange(t *testing.T) {
	assert.Any(t, Range{L: 1, R: 2}, New(1, 2))
}
