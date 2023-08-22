package hypertext

import (
	"bufio"
	"github.com/PuerkitoBio/goquery"
	assert2 "github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestTitle(t *testing.T) {
	f, openFail := fixtureFile("test.html")
	assert2.FatalOnError(t, openFail)

	d, readFail := goquery.NewDocumentFromReader(bufio.NewReader(f))
	assert2.FatalOnError(t, readFail)

	assert2.String(t, "Test Title", Title(d))
}
