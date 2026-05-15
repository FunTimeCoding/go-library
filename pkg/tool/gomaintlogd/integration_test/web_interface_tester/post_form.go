package web_interface_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (o *Tester) PostForm(
	path string,
	values url.Values,
) string {
	o.t.Helper()
	r, e := http.Post(
		join.Empty(o.base, path),
		constant.FormEncoded,
		strings.NewReader(values.Encode()),
	)
	assert.FatalOnError(o.t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(o.t, f)

	return string(b)
}
