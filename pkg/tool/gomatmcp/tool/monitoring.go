package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) Monitoring(
	_ context.Context,
	_ mcp.CallToolRequest,
) (result *mcp.CallToolResult, e error) {
	defer func() {
		if r := recover(); r != nil {
			result = mcp.NewToolResultError(fmt.Sprintf("%v", r))
			e = nil
		}
	}()

	if f := t.monitor.RunNow(); f != nil {
		return response.Fail("monitoring failed: %v", f)
	}

	return response.SuccessAny(
		map[string]any{
			"message": "Topic monitoring executed successfully",
		},
	)
}
