package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestReduceSpace(t *testing.T) {
	assert.String(t, " ", ReduceSpace("  "))
}
