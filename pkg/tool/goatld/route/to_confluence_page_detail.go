package route

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
)

func toConfluencePageDetail(p *page.Page) server.ConfluencePageDetail {
	result := server.ConfluencePageDetail{
		Identifier: p.Identifier,
		Title:      p.Name,
	}

	if p.Link != "" {
		result.Link = &p.Link
	}

	if p.Raw.SpaceId != "" {
		result.SpaceIdentifier = &p.Raw.SpaceId
	}

	if p.Raw.Body.Storage.Value != "" {
		body := page.ToMarkdown(p.Raw.Body.Storage.Value)
		result.Body = &body
	}

	return result
}
