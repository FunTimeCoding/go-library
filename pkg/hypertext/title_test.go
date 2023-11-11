package hypertext

import (
	"bufio"
	"github.com/PuerkitoBio/goquery"
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestTitle(t *testing.T) {
	f, openFail := fixtureFile("test.html")
	assert.FatalOnError(t, openFail)

	d, readFail := goquery.NewDocumentFromReader(bufio.NewReader(f))
	assert.FatalOnError(t, readFail)

	assert.String(t, "Test Title", Title(d))
}
