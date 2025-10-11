package page

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"

func New(
	p *response.Page,
	host string,
) *Page {
	return &Page{
		Identifier: p.Id,
		Name:       p.Title,
		Link:       link(p, host, false),
		TinyLink:   link(p, host, true),
		Raw:        p,
	}
}
