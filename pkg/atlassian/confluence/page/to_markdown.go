package page

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func ToMarkdown(markup string) string {
	c := md.NewConverter("", true, nil)
	result, e := c.ConvertString(markup)
	errors.PanicOnError(e)

	return result
}
