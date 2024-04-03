package runtime

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	assert.True(t, Version() != nil)
}
