package gomaintlogd

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"testing"
)

func TestModelContextClient(t *testing.T) {
	s, port, l := setup(t)
	defer s.Close()
	defer l.Stop()
	x := context.Background()
	c := mcp.NewClient(
		&mcp.Implementation{
			Name:    "test-client",
			Version: constant.DefaultVersion,
		},
		nil,
	)
	endpoint := fmt.Sprintf("http://localhost:%d/mcp", port)
	session, e := c.Connect(
		x,
		&mcp.StreamableClientTransport{Endpoint: endpoint},
		nil,
	)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(session)
	// ListTools
	tools, e := session.ListTools(x, &mcp.ListToolsParams{})
	assert.FatalOnError(t, e)
	assert.Count(t, 4, tools.Tools)
	// add_entry
	add, e := session.CallTool(
		x,
		&mcp.CallToolParams{
			Name: "add_entry",
			Arguments: map[string]any{
				"action":      "restarted web server",
				"user":        "jdoe",
				"system":      "worker1",
				"service":     "nginx",
				"description": "nginx was unresponsive",
			},
		},
	)
	assert.FatalOnError(t, e)
	assert.False(t, add.IsError)
	// list_entries
	list, e := session.CallTool(
		x,
		&mcp.CallToolParams{
			Name:      "list_entries",
			Arguments: map[string]any{},
		},
	)
	assert.FatalOnError(t, e)
	assert.False(t, list.IsError)
	assert.StringContains(
		t,
		"1 entries",
		list.Content[0].(*mcp.TextContent).Text,
	)
	// list_entries - filter by system
	filtered, e := session.CallTool(
		x,
		&mcp.CallToolParams{
			Name:      "list_entries",
			Arguments: map[string]any{"system": "worker1"},
		},
	)
	assert.FatalOnError(t, e)
	assert.False(t, filtered.IsError)
	assert.StringContains(
		t,
		"1 entries",
		filtered.Content[0].(*mcp.TextContent).Text,
	)
	// list_entries - filter by user
	filteredUser, e := session.CallTool(
		x,
		&mcp.CallToolParams{
			Name:      "list_entries",
			Arguments: map[string]any{"user": "jdoe"},
		},
	)
	assert.FatalOnError(t, e)
	assert.False(t, filteredUser.IsError)
	assert.StringContains(
		t,
		"1 entries",
		filteredUser.Content[0].(*mcp.TextContent).Text,
	)
	// list_entries - filter no match
	noMatch, e := session.CallTool(
		x,
		&mcp.CallToolParams{
			Name:      "list_entries",
			Arguments: map[string]any{"system": "nonexistent"},
		},
	)
	assert.FatalOnError(t, e)
	assert.False(t, noMatch.IsError)
	assert.StringContains(
		t,
		"0 entries",
		noMatch.Content[0].(*mcp.TextContent).Text,
	)
	// update_entry
	update, e := session.CallTool(
		x,
		&mcp.CallToolParams{
			Name: "update_entry",
			Arguments: map[string]any{
				"id":     float64(1),
				"action": "cleared and documented",
			},
		},
	)
	assert.FatalOnError(t, e)
	assert.False(t, update.IsError)
	assert.StringContains(
		t,
		"updated successfully",
		update.Content[0].(*mcp.TextContent).Text,
	)
	// list_entries - verify update
	list, e = session.CallTool(
		x,
		&mcp.CallToolParams{
			Name:      "list_entries",
			Arguments: map[string]any{},
		},
	)
	assert.FatalOnError(t, e)
	assert.StringContains(
		t,
		"cleared and documented",
		list.Content[0].(*mcp.TextContent).Text,
	)
	// delete_entry
	del, e := session.CallTool(
		x,
		&mcp.CallToolParams{
			Name:      "delete_entry",
			Arguments: map[string]any{"id": float64(1)},
		},
	)
	assert.FatalOnError(t, e)
	assert.False(t, del.IsError)
	assert.StringContains(
		t,
		"deleted successfully",
		del.Content[0].(*mcp.TextContent).Text,
	)
	// list_entries - verify delete
	list, e = session.CallTool(
		x,
		&mcp.CallToolParams{
			Name:      "list_entries",
			Arguments: map[string]any{},
		},
	)
	assert.FatalOnError(t, e)
	assert.StringContains(
		t,
		"0 entries",
		list.Content[0].(*mcp.TextContent).Text,
	)
}
