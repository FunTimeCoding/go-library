package web_interface_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"io"
	"net/http"
)

func (o *Tester) Get(path string) string {
	o.t.Helper()
	r, e := http.Get(join.Empty(o.base, path))
	assert.FatalOnError(o.t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(o.t, f)

	return string(b)
}
