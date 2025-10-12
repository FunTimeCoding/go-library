package page

import (
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func ToMarkdown(markup string) string {
	c := converter.NewConverter()
	result, e := c.ConvertString(markup)
	errors.PanicOnError(e)

	return result
}
