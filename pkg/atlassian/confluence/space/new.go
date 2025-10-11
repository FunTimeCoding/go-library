package space

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"

func New(p *response.Space) *Space {
	return &Space{Identifier: p.Id, Name: p.Name, Raw: p}
}
