package route

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/search_result"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
)

func toConfluencePages(results []*search_result.Result) []server.ConfluencePage {
	pages := make([]server.ConfluencePage, 0, len(results))

	for _, r := range results {
		entry := server.ConfluencePage{
			Identifier: r.Raw.Id,
			Title:      r.Name,
		}

		if r.Raw.Type != "" {
			entry.Type = &r.Raw.Type
		}

		if r.Raw.Status != "" {
			entry.Status = &r.Raw.Status
		}

		if r.Raw.Links.Self != "" {
			entry.Link = &r.Raw.Links.Self
		}

		pages = append(pages, entry)
	}

	return pages
}
