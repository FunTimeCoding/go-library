package detail

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestPrometheus(t *testing.T) {
	assert.True(t, New() != nil)
}
