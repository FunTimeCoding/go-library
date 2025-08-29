package thread

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestThread(t *testing.T) {
	assert.NotNil(t, New(nil, nil))
}
