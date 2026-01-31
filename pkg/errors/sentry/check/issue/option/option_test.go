package option

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestOption(t *testing.T) {
	assert.NotNil(t, New())
}
