package gomatmcp

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/tool"
)

func addTool(
	s *server.MCPServer,
	t *tool.Tool,
	withMonitor bool,
) {
	s.AddTool(
		mcp.NewTool(
			constant.ListChannels,
			mcp.WithDescription("List public channels in the Mattermost workspace"),
			mcp.WithNumber(
				"limit",
				mcp.Description("Max channels to return (default 100)"),
				mcp.DefaultNumber(100),
			),
			mcp.WithNumber(
				"page",
				mcp.Description("Page number starting from 0"),
				mcp.DefaultNumber(0),
			),
		),
		mcp.NewTypedToolHandler(t.ListChannels),
	)
	s.AddTool(
		mcp.NewTool(
			constant.GetChannelHistory,
			mcp.WithDescription(
				"Get recent messages from a channel. Provide channel_id or channel_name.",
			),
			mcp.WithString(
				"channel_id",
				mcp.Description("Channel ID"),
			),
			mcp.WithString(
				"channel_name",
				mcp.Description("Channel name"),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Number of messages (default 30)"),
				mcp.DefaultNumber(30),
			),
		),
		mcp.NewTypedToolHandler(t.GetChannelHistory),
	)
	s.AddTool(
		mcp.NewTool(
			constant.PostMessage,
			mcp.WithDescription(
				"Post a message to a Mattermost channel",
			),
			mcp.WithString(
				"channel_id",
				mcp.Description("Channel ID"),
			),
			mcp.WithString(
				"channel_name",
				mcp.Description("Channel name"),
			),
			mcp.WithString(
				"message",
				mcp.Required(),
				mcp.Description("Message text"),
			),
		),
		mcp.NewTypedToolHandler(t.PostMessage),
	)
	s.AddTool(
		mcp.NewTool(
			constant.ReplyToThread,
			mcp.WithDescription("Reply to a message thread"),
			mcp.WithString(
				"channel_id",
				mcp.Required(),
				mcp.Description("Channel ID"),
			),
			mcp.WithString(
				"post_id",
				mcp.Required(),
				mcp.Description("Parent post ID"),
			),
			mcp.WithString(
				"message",
				mcp.Required(),
				mcp.Description("Reply text"),
			),
		),
		mcp.NewTypedToolHandler(t.ReplyToThread),
	)
	s.AddTool(
		mcp.NewTool(
			constant.AddReaction,
			mcp.WithDescription("Add an emoji reaction to a post"),
			mcp.WithString(
				"post_id",
				mcp.Required(),
				mcp.Description("Post ID"),
			),
			mcp.WithString(
				"emoji_name",
				mcp.Required(),
				mcp.Description("Emoji name without colons"),
			),
		),
		mcp.NewTypedToolHandler(t.AddReaction),
	)
	s.AddTool(
		mcp.NewTool(
			constant.GetThreadReplies,
			mcp.WithDescription(
				"Get all replies in a message thread",
			),
			mcp.WithString(
				"post_id",
				mcp.Required(),
				mcp.Description("Parent post ID"),
			),
		),
		mcp.NewTypedToolHandler(t.GetThreadReplies),
	)
	s.AddTool(
		mcp.NewTool(
			constant.GetUsers,
			mcp.WithDescription(
				"List users in the Mattermost workspace",
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Max users to return (default 100)"),
				mcp.DefaultNumber(100),
			),
			mcp.WithNumber(
				"page",
				mcp.Description("Page number starting from 0"),
				mcp.DefaultNumber(0),
			),
		),
		mcp.NewTypedToolHandler(t.GetUsers),
	)
	s.AddTool(
		mcp.NewTool(
			constant.GetUserProfile,
			mcp.WithDescription(
				"Get profile information for a user",
			),
			mcp.WithString(
				"user_id",
				mcp.Required(),
				mcp.Description("User ID"),
			),
		),
		mcp.NewTypedToolHandler(t.GetUserProfile),
	)

	if withMonitor {
		s.AddTool(
			mcp.NewTool(
				constant.RunMonitoring,
				mcp.WithDescription(
					"Run topic monitoring immediately",
				),
			),
			t.RunMonitoring,
		)
	}
}
