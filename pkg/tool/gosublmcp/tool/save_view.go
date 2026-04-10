package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/request"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type saveViewArguments struct {
	ViewIdentifier request.Integer `json:"view_id"`
	FilePath       string `json:"file_path"`
}

func (t *Tool) SaveView(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments saveViewArguments,
) (*mcp.CallToolResult, error) {
	e := t.client.SaveView(int(arguments.ViewIdentifier), arguments.FilePath)

	if e != nil {
		return response.Fail("save view: %v", e)
	}

	return response.Success("saved")
}
