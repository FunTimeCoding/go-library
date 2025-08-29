package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestWithVerbose(t *testing.T) {
	assert.NotNil(t, WithVerbose(true))
}
