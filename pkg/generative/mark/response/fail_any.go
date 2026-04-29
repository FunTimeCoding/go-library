package response

import (
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func FailAny(v any) (*mcp.CallToolResult, error) {
	return Fail(notation.MarshalIndent(v))
}
