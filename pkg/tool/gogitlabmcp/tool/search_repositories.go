package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

type searchRepositoriesArguments struct {
	Query string `json:"query"`
}

func (t *Tool) SearchRepositories(
	_ context.Context,
	_ mcp.CallToolRequest,
	a searchRepositoriesArguments,
) (*mcp.CallToolResult, error) {
	if a.Query == "" {
		return response.Fail("query is required")
	}

	v, _, e := t.client.Search.Projects(
		a.Query,
		&gitlab.SearchOptions{ListOptions: gitlab.ListOptions{PerPage: 20}},
	)

	if e != nil {
		return response.Fail("search repositories: %v", e)
	}

	return response.SuccessAny(v)
}
