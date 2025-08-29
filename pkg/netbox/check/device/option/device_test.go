package option

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestDevice(t *testing.T) {
	assert.NotNil(t, New())
}
