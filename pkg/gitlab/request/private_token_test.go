package request

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"testing"
)

func TestPrivateToken(t *testing.T) {
	r := web.NewGet("http://localhost")
	PrivateToken(r, "test")
	assert.Any(t, http.Header{"Private-Token": {"test"}}, r.Header)
}
