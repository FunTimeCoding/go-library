package convert

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/generated/server"
)

func ConfluencePageDetail(p *page.Page) *server.ConfluencePageDetail {
	result := &server.ConfluencePageDetail{
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
		result.Body = new(page.ToMarkdown(p.Raw.Body.Storage.Value))
	}

	return result
}
