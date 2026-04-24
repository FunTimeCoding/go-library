package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goitermmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) SendText(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SendText,
) (*mcp.CallToolResult, error) {
	if a.Text == "" {
		return response.Fail("text is required")
	}

	e := t.client.SendText(a.SessionIdentifier, a.Text)

	if e != nil {
		return response.Fail("send text: %v", e)
	}

	if a.SendEnter {
		f := t.client.SendKey(a.SessionIdentifier, "enter")

		if f != nil {
			return response.Fail("send enter: %v", f)
		}
	}

	return response.Success("sent: %s", a.Text)
}
