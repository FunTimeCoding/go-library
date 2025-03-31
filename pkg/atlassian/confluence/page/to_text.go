package page

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"golang.org/x/net/html"
	"strings"
)

func ToText(markup string) string {
	n, e := html.Parse(strings.NewReader(markup))
	errors.PanicOnError(e)
	var b strings.Builder
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			b.WriteString(" " + n.Data)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)

	return b.String()
}
