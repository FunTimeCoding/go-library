package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/request"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type setTabColorArguments struct {
	SessionIdentifier string          `json:"session_id"`
	Red               request.Integer `json:"red"`
	Green             request.Integer `json:"green"`
	Blue              request.Integer `json:"blue"`
}

func (t *Tool) SetTabColor(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments setTabColorArguments,
) (*mcp.CallToolResult, error) {
	e := t.client.SetTabColor(
		arguments.SessionIdentifier,
		int(arguments.Red),
		int(arguments.Green),
		int(arguments.Blue),
	)

	if e != nil {
		return response.Fail("set tab color: %v", e)
	}

	return response.Success("color set")
}
