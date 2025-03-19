package advanced_parameter

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestParameter(t *testing.T) {
	assert.True(t, New() != nil)
}
