package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/request"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

type sendKeyArguments struct {
	SessionIdentifier string          `json:"session_id"`
	Keys              []string        `json:"keys"`
	Interval          request.Integer `json:"interval"`
}

func (t *Tool) SendKey(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments sendKeyArguments,
) (*mcp.CallToolResult, error) {
	if len(arguments.Keys) == 0 {
		return response.Fail("keys is required")
	}

	interval := time.Duration(arguments.Interval) * time.Second

	if interval < time.Second {
		interval = time.Second
	}

	for i, key := range arguments.Keys {
		if i > 0 {
			time.Sleep(interval)
		}

		e := t.client.SendKey(
			arguments.SessionIdentifier,
			key,
		)

		if e != nil {
			return response.Fail(
				"send key %s (step %d): %v",
				key,
				i+1,
				e,
			)
		}
	}

	return response.Success(
		"sent keys: %s",
		join.Comma(arguments.Keys),
	)
}
