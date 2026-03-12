package mock_listener

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestMockListener(t *testing.T) {
	assert.NotNil(t, New())
}
