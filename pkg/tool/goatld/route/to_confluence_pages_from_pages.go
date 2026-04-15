package route

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
)

func toConfluencePagesFromPages(pages []*page.Page) []server.ConfluencePage {
	result := make([]server.ConfluencePage, 0, len(pages))

	for _, p := range pages {
		entry := server.ConfluencePage{
			Identifier: p.Identifier,
			Title:      p.Name,
		}

		if p.Link != "" {
			entry.Link = &p.Link
		}

		result = append(result, entry)
	}

	return result
}
