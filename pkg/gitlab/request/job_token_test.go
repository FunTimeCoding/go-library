package request

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"testing"
)

func TestJobToken(t *testing.T) {
	r := web.NewGet("http://localhost")
	JobToken(r, "test")
	assert.Any(t, http.Header{"Job-Token": {"test"}}, r.Header)
}
