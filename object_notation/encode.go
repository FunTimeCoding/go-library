package object_notation

import (
	"bytes"
	"encoding/json"
	"github.com/funtimecoding/go-library/errors"
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

	errors.FatalOnError(encoder.Encode(object))

	return strings.TrimSuffix(buffer.String(), "\n")
}
