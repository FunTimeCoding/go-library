package request_context

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestContext(t *testing.T) {
	assert.True(t, New(nil, nil) != nil)
}
