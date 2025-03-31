package space

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"

type Space struct {
	Name string

	Raw *response.Space
}
