package example

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func Server() {
	s := server.NewMCPServer(
		"Demo",
		"1.0.0",
		server.WithToolCapabilities(false),
	)
	s.AddTool(
		mcp.NewTool(
			"hello_world",
			mcp.WithDescription("Say hello to someone"),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Name of the person to greet"),
			),
		),
		hello,
	)
	errors.PanicOnError(server.ServeStdio(s))
}

func hello(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, e := r.RequireString("name")

	if e != nil {
		return mcp.NewToolResultError(e.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Hello %s", name)), nil
}
