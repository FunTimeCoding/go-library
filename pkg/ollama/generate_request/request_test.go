package generate_request

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestRequest(t *testing.T) {
	assert.True(t, New(strings.Alfa) != nil)
}
