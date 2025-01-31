package web_response

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestResponse(t *testing.T) {
	assert.True(t, New(nil) != nil)
}
