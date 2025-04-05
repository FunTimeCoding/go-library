package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"net/http"
	"testing"
)

func TestBearer(t *testing.T) {
	r := NewGet(Localhost)
	Bearer(r, strings.Alfa)
	assert.Any(
		t,
		http.Header{"Authorization": {"Bearer Alfa"}},
		r.Header,
	)
}
