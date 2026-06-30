package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
	"testing"
)

func TestToken(t *testing.T) {
	r := NewGet(constant.Localhost)
	Token(r, upper.Alfa)
	assert.Any(
		t,
		http.Header{"Authorization": {"Token Alfa"}},
		r.Header,
	)
}
