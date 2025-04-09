package page

import (
	"bytes"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/yuin/goldmark"
)

func ToMarkdown(markup string) string {
	c := md.NewConverter("", true, nil)
	result, e := c.ConvertString(markup)
	errors.PanicOnError(e)

	return result
}

func ToMarkup(markdown string) string {
	var b bytes.Buffer

	errors.PanicOnError(goldmark.Convert([]byte(markdown), &b))

	return b.String()
}
