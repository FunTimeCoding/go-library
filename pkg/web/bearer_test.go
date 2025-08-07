package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
	"testing"
)

func TestBearer(t *testing.T) {
	r := NewGet(constant.Localhost)
	Bearer(r, strings.Alfa)
	assert.Any(
		t,
		http.Header{"Authorization": {"Bearer Alfa"}},
		r.Header,
	)
}
