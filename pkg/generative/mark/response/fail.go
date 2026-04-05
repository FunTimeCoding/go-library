package response

import (
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

func Fail(
	f string,
	a ...any,
) (*mcp.CallToolResult, error) {
	if len(a) > 0 {
		f = fmt.Sprintf(f, a...)
	}

	return mcp.NewToolResultError(f), nil
}
