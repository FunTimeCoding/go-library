package generate_request

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestRequest(t *testing.T) {
	assert.NotNil(t, New())
}
