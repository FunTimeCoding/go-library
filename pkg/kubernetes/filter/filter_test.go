package filter

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	assert.NotNil(t, New())
}
