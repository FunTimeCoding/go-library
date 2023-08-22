package object_notation

import (
	"bytes"
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func Encode(
	object any,
	indent bool,
) string {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)

	if indent {
		encoder.SetIndent("", "    ")
	}

	errors.PanicOnError(encoder.Encode(object))

	return strings.TrimSuffix(buffer.String(), separator.Unix)
}
