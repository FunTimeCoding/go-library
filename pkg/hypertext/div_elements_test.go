package hypertext

import (
	"bufio"
	"github.com/PuerkitoBio/goquery"
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestDivElements(t *testing.T) {
	f, openFail := fixtureFile("test.html")
	assert.FatalOnError(t, openFail)

	d, readFail := goquery.NewDocumentFromReader(bufio.NewReader(f))
	assert.FatalOnError(t, readFail)

	assert.Any(t, []string{"Example DT", "Example DD"}, DivElements(d))
}
