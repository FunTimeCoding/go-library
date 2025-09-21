package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
	"testing"
)

func TestToken(t *testing.T) {
	r := NewGet(constant.Localhost)
	Token(r, strings.Alfa)
	assert.Any(
		t,
		http.Header{"Authorization": {"Token Alfa"}},
		r.Header,
	)
}
