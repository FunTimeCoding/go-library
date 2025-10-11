package request

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"net/http"
	"testing"
)

func TestJobToken(t *testing.T) {
	r := web.NewGet(locator.New(constant.Localhost).Insecure().String())
	JobToken(r, "test")
	assert.Any(t, http.Header{"Job-Token": {"test"}}, r.Header)
}
