package page

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"

func New(
	p *response.Page,
	host string,
) *Page {
	return &Page{
		Name: p.Title,
		Raw:  p,
	}
}
