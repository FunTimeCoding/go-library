package request

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/web"
	"testing"
)

func TestJobToken(t *testing.T) {
	r := web.NewGet("http://localhost")
	JobToken(r, "test")
	assert.Any(t, map[string][]string{"Job-Token": {"test"}}, r.Header)
}
