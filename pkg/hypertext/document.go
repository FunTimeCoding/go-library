package hypertext

import (
	"bufio"
	"github.com/PuerkitoBio/goquery"
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
)

func Document(f io.Reader) *goquery.Document {
	result, e := goquery.NewDocumentFromReader(bufio.NewReader(f))
	errors.PanicOnError(e)

	return result
}
