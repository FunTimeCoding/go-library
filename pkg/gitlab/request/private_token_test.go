package request

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/web"
	"testing"
)

func TestPrivateToken(t *testing.T) {
	r := web.NewGet("http://localhost")
	PrivateToken(r, "test")
	assert.Any(t, map[string][]string{"Private-Token": {"test"}}, r.Header)
}
