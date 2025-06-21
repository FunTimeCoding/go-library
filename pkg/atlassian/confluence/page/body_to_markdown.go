package page

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"
	"github.com/funtimecoding/go-library/pkg/text"
	"github.com/funtimecoding/go-library/pkg/text/option"
)

func bodyToMarkdown(b response.Body) string {
	if b.Storage.Value == "" {
		return ""
	}

	result := ToMarkdown(b.Storage.Value)
	p := option.New()
	p.AllowedBlankLines = 0
	p.NewlineAtEnd = false
	result = text.OptimizeWhitespace(result, p)

	return result
}
