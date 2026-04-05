package response

import (
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

func Success(
	f string,
	a ...any,
) (*mcp.CallToolResult, error) {
	if len(a) > 0 {
		f = fmt.Sprintf(f, a...)
	}

	return mcp.NewToolResultText(f), nil
}
