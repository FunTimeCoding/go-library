package example

import (
	"fmt"
	"github.com/JohannesKaufmann/html-to-markdown"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/text"
	"github.com/funtimecoding/go-library/pkg/text/option"
	"golang.org/x/net/html"
	"strings"
)

func Page() {
	c := confluence.NewEnvironment()

	if true {
		fmt.Println("PagesBasic")

		for _, p := range c.PagesBasic() {
			fmt.Printf("  Page: %s\n", p.Title)

			if false {
				if p.Title != "Example" {
					continue
				}
			}

			if false {
				fmt.Printf("    Raw: %+v\n", p)
			}

			printBody(p.Body)
		}
	}
}

func printBody(b response.Body) {
	if false {
		fmt.Printf("    Storage: %s\n", b.Storage.Value)
	}

	if false {
		fmt.Printf("    Text: %s\n", TextFromMarkup(b.Storage.Value))
	}

	markdown := ToMarkdown(b.Storage.Value)
	p := option.New()
	p.AllowedBlankLines = 0
	p.NewlineAtEnd = false
	markdown = text.OptimizeWhitespace(markdown, p)
	fmt.Printf("    Markdown: %s\n", markdown)
}

func TextFromMarkup(markup string) string {
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

func ToMarkdown(markup string) string {
	c := md.NewConverter("", true, nil)
	result, e := c.ConvertString(markup)
	errors.PanicOnError(e)

	return result
}
