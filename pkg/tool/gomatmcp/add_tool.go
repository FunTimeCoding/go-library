package gomatmcp

import (
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/tool"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func addTool(
	s *server.MCPServer,
	t *tool.Tool,
	withMonitor bool,
) {
	s.AddTool(
		mcp.NewTool(
			constant.MyChannels,
			mcp.WithDescription(
				"List your channels sorted by most recent activity. Includes public, private, DM, and group DM channels.",
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Max channels to return (default 30)"),
				mcp.DefaultNumber(30),
			),
			mcp.WithString(
				"type",
				mcp.Description(
					"Filter by type: D (DM), G (group DM), O (public), P (private). Omit for all.",
				),
			),
			mcp.WithString(
				"since",
				mcp.Description(
					"Only channels with activity since this time (e.g. \"2026-04-22 15:00\")",
				),
			),
		),
		mcp.NewTypedToolHandler(t.MyChannels),
	)
	s.AddTool(
		mcp.NewTool(
			constant.MyThreads,
			mcp.WithDescription(
				"List your subscribed threads sorted by most recent reply. Shows unread counts, participants, and the original message.",
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Max threads to return (default 20)"),
				mcp.DefaultNumber(20),
			),
			mcp.WithBoolean(
				"unread_only",
				mcp.Description(
					"Only show threads with unread replies",
				),
			),
			mcp.WithString(
				"since",
				mcp.Description(
					"Only threads with activity since this time (e.g. \"2026-04-22 15:00\")",
				),
			),
		),
		mcp.NewTypedToolHandler(t.MyThreads),
	)
	s.AddTool(
		mcp.NewTool(
			constant.SearchChannels,
			mcp.WithDescription(
				"Search for channels by name. Returns channels you have access to.",
			),
			mcp.WithString(
				"query",
				mcp.Required(),
				mcp.Description("Channel name to search for"),
			),
		),
		mcp.NewTypedToolHandler(t.SearchChannels),
	)
	s.AddTool(
		mcp.NewTool(
			constant.SearchMessages,
			mcp.WithDescription(
				"Search messages across all channels. Supports Mattermost search syntax: from:username, in:channelname, before:2026-04-22, after:2026-04-20. Plain text matches posts containing those words.",
			),
			mcp.WithString(
				"terms",
				mcp.Required(),
				mcp.Description("Search terms"),
			),
		),
		mcp.NewTypedToolHandler(t.SearchMessages),
	)
	s.AddTool(
		mcp.NewTool(
			constant.SearchUsers,
			mcp.WithDescription(
				"Search users by username, first name, last name, or nickname.",
			),
			mcp.WithString(
				"query",
				mcp.Required(),
				mcp.Description("Search term"),
			),
		),
		mcp.NewTypedToolHandler(t.SearchUsers),
	)
	s.AddTool(
		mcp.NewTool(
			constant.DownloadFile,
			mcp.WithDescription(
				"Download a file from Mattermost to a local path. Defaults to /tmp/<filename>.",
			),
			mcp.WithString(
				"file_id",
				mcp.Required(),
				mcp.Description("File ID from a post's file_ids"),
			),
			mcp.WithString(
				"path",
				mcp.Description(
					"Local path to save to (default: /tmp/<filename>)",
				),
			),
		),
		mcp.NewTypedToolHandler(t.DownloadFile),
	)
	s.AddTool(
		mcp.NewTool(
			constant.UploadFile,
			mcp.WithDescription(
				"Upload a local file to a Mattermost channel or DM as an attachment.",
			),
			mcp.WithString(
				"path",
				mcp.Required(),
				mcp.Description("Local file path to upload"),
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
				mcp.Description(
					"Message to post with the file (default: filename)",
				),
			),
		),
		mcp.NewTypedToolHandler(t.UploadFile),
	)
	s.AddTool(
		mcp.NewTool(
			constant.SendDirectMessage,
			mcp.WithDescription(
				"Send a direct message to a user by username.",
			),
			mcp.WithString(
				"username",
				mcp.Required(),
				mcp.Description("Recipient's username"),
			),
			mcp.WithString(
				"message",
				mcp.Required(),
				mcp.Description("Message text"),
			),
		),
		mcp.NewTypedToolHandler(t.SendDirectMessage),
	)
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
			mcp.WithString(
				"since",
				mcp.Description(
					"Only messages since this time (e.g. \"2026-04-22 15:00\"). Overrides limit.",
				),
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
			t.Monitoring,
		)
	}
}
