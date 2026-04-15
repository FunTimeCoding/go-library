package page_put

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"

type Put struct {
	Identifier string           `json:"id"`
	Status     string           `json:"status"`
	Title      string           `json:"title"`
	Body       response.Storage `json:"body"`
	Version    Version          `json:"version"`
}
