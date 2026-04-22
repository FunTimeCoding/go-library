package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type sendTextArguments struct {
	SessionIdentifier string `json:"session_id"`
	Text              string `json:"text"`
	SendEnter         bool   `json:"send_enter"`
}

func (t *Tool) SendText(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments sendTextArguments,
) (*mcp.CallToolResult, error) {
	if arguments.Text == "" {
		return response.Fail("text is required")
	}

	e := t.client.SendText(arguments.SessionIdentifier, arguments.Text)

	if e != nil {
		return response.Fail("send text: %v", e)
	}

	if arguments.SendEnter {
		f := t.client.SendKey(arguments.SessionIdentifier, "enter")

		if f != nil {
			return response.Fail("send enter: %v", f)
		}
	}

	return response.Success("sent: %s", arguments.Text)
}
