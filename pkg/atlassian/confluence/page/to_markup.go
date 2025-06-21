package page

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/yuin/goldmark"
)

func ToMarkup(markdown string) string {
	var b bytes.Buffer

	errors.PanicOnError(goldmark.Convert([]byte(markdown), &b))

	return b.String()
}
