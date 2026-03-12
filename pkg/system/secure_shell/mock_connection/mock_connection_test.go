package mock_connection

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestMockConnection(t *testing.T) {
	assert.NotNil(t, New())
}
