package web_response

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestResponse(t *testing.T) {
	assert.NotNil(t, New(nil))
}
