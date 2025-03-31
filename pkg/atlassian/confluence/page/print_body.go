package page

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"
	"github.com/funtimecoding/go-library/pkg/text"
	"github.com/funtimecoding/go-library/pkg/text/option"
)

func PrintBody(b response.Body) {
	if false {
		fmt.Printf("    Storage: %s\n", b.Storage.Value)
	}

	if false {
		fmt.Printf("    Text: %s\n", ToText(b.Storage.Value))
	}

	markdown := ToMarkdown(b.Storage.Value)
	p := option.New()
	p.AllowedBlankLines = 0
	p.NewlineAtEnd = false
	markdown = text.OptimizeWhitespace(markdown, p)
	fmt.Printf("    Markdown: %s\n", markdown)
}
