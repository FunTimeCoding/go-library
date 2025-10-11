package page

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"

type Page struct {
	Identifier string
	Name       string
	Link       string
	TinyLink   string

	Raw *response.Page
}
