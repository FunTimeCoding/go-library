package chat_request

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestRequest(t *testing.T) {
	assert.True(t, New() != nil)
}
