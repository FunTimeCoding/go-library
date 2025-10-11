package space

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"

type Space struct {
	Identifier string
	Name       string
	Raw        *response.Space
}
