package callback

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestServer(t *testing.T) {
	assert.True(t, New(0, false) != nil)
}
