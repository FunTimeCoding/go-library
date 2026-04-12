package gosublmcp

import (
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/tool"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func addTool(
	s *server.MCPServer,
	t *tool.Tool,
) {
	s.AddTool(
		mcp.NewTool(
			constant.ListViews,
			mcp.WithDescription(
				"List all open views across Sublime Text windows. "+
					"View IDs are session-stable (persist across plugin reloads). "+
					"Unsaved scratch buffers survive Sublime restarts.",
			),
		),
		mcp.NewTypedToolHandler(t.ListViews),
	)
	s.AddTool(
		mcp.NewTool(
			constant.ReadView,
			mcp.WithDescription(
				"Read full text and metadata of a Sublime Text view",
			),
			mcp.WithNumber(
				"view_id",
				mcp.Required(),
				mcp.Description("Buffer identifier (stable within session)"),
			),
		),
		mcp.NewTypedToolHandler(t.ReadView),
	)
	s.AddTool(
		mcp.NewTool(
			constant.EditView,
			mcp.WithDescription(
				"Line-based text replacement in a Sublime Text view",
			),
			mcp.WithNumber(
				"view_id",
				mcp.Required(),
				mcp.Description("Buffer identifier (stable within session)"),
			),
			mcp.WithString(
				"old_string",
				mcp.Required(),
				mcp.Description("Text to find and replace"),
			),
			mcp.WithString(
				"new_string",
				mcp.Required(),
				mcp.Description("Replacement text"),
			),
			mcp.WithBoolean(
				"replace_all",
				mcp.Description(
					"Replace all occurrences instead of requiring unique match",
				),
			),
		),
		mcp.NewTypedToolHandler(t.EditView),
	)
	s.AddTool(
		mcp.NewTool(
			constant.CreateView,
			mcp.WithDescription(
				"Create a new scratch buffer in Sublime Text",
			),
			mcp.WithString(
				"title",
				mcp.Description("Buffer title"),
			),
			mcp.WithString(
				"content",
				mcp.Description("Initial buffer content"),
			),
			mcp.WithString(
				"syntax",
				mcp.Description("Syntax name (e.g. Go, Python, Markdown)"),
			),
		),
		mcp.NewTypedToolHandler(t.CreateView),
	)
	s.AddTool(
		mcp.NewTool(
			constant.SaveView,
			mcp.WithDescription(
				"Save a Sublime Text view to disk",
			),
			mcp.WithNumber(
				"view_id",
				mcp.Required(),
				mcp.Description("Buffer identifier (stable within session)"),
			),
			mcp.WithString(
				"file_path",
				mcp.Description(
					"File path to save to (required for unsaved buffers)",
				),
			),
		),
		mcp.NewTypedToolHandler(t.SaveView),
	)
	s.AddTool(
		mcp.NewTool(
			constant.CloseView,
			mcp.WithDescription("Close a Sublime Text tab"),
			mcp.WithNumber(
				"view_id",
				mcp.Required(),
				mcp.Description("Buffer identifier (stable within session)"),
			),
		),
		mcp.NewTypedToolHandler(t.CloseView),
	)
	s.AddTool(
		mcp.NewTool(
			constant.OpenFile,
			mcp.WithDescription(
				"Open a file from disk as a Sublime Text tab",
			),
			mcp.WithString(
				"file_path",
				mcp.Required(),
				mcp.Description("Absolute path to the file"),
			),
		),
		mcp.NewTypedToolHandler(t.OpenFile),
	)
}
