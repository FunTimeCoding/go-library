package notation

import (
	"bytes"
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func Encode(
	a any,
	indent bool,
) string {
	b := &bytes.Buffer{}
	e := json.NewEncoder(b)
	e.SetEscapeHTML(false)

	if indent {
		e.SetIndent("", "    ")
	}

	errors.PanicOnError(e.Encode(a))

	return strings.TrimSuffix(b.String(), separator.Unix)
}
