package thread

import (
	"testing"

	"github.com/funtimecoding/go-library/pkg/assert"
)

func TestThread(t *testing.T) {
	assert.True(t, New(nil, nil) != nil)
}
