package request

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"net/http"
	"testing"
)

func TestPrivateToken(t *testing.T) {
	r := web.NewGet(locator.New(constant.Localhost).Insecure().String())
	PrivateToken(r, "test")
	assert.Any(t, http.Header{"Private-Token": {"test"}}, r.Header)
}
