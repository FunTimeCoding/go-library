package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
)

type myThreadsArguments struct {
	Limit      int    `json:"limit"`
	UnreadOnly bool   `json:"unread_only"`
	Since      string `json:"since"`
}

func (t *Tool) MyThreads(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments myThreadsArguments,
) (result *mcp.CallToolResult, e error) {
	defer func() {
		if r := recover(); r != nil {
			result = mcp.NewToolResultError(fmt.Sprintf("%v", r))
			e = nil
		}
	}()
	limit := arguments.Limit

	if limit <= 0 {
		limit = 20
	}

	me := t.client.Me()
	team := t.client.DefaultTeam()
	opts := model.GetUserThreadsOpts{
		PageSize: uint64(limit),
		Extended: true,
	}

	if arguments.Since != "" {
		since, g := parseSince(arguments.Since)

		if g != nil {
			return response.Fail(
				"invalid since format: %v",
				g,
			)
		}

		opts.Since = uint64(since.UnixMilli())
	}

	threads, _, f := t.client.Nested().GetUserThreads(
		t.client.Context(),
		me.Id,
		team.Id,
		opts,
	)

	if f != nil {
		return response.Fail("get threads failed: %v", f)
	}

	type participant struct {
		Username string `json:"username"`
	}
	type row struct {
		PostIdentifier string        `json:"post_identifier"`
		Channel        string        `json:"channel,omitempty"`
		Author         string        `json:"author"`
		Message        string        `json:"message"`
		ReplyCount     int64         `json:"reply_count"`
		UnreadReplies  int64         `json:"unread_replies"`
		UnreadMentions int64         `json:"unread_mentions"`
		LastReplyAt    string        `json:"last_reply_at"`
		Participants   []participant `json:"participants,omitempty"`
	}
	var rows []row

	for _, thread := range threads.Threads {
		if arguments.UnreadOnly && thread.UnreadReplies == 0 {
			continue
		}

		r := row{
			PostIdentifier: thread.PostId,
			ReplyCount:     thread.ReplyCount,
			UnreadReplies:  thread.UnreadReplies,
			UnreadMentions: thread.UnreadMentions,
			LastReplyAt:    formatMilli(thread.LastReplyAt),
		}

		if thread.Post != nil {
			r.Message = thread.Post.Message

			if thread.Post.UserId != "" {
				author := t.client.User(thread.Post.UserId)
				r.Author = author.Username
			}

			ch := t.client.Channel(thread.Post.ChannelId)
			r.Channel = t.channelDisplayName(ch)
		}

		for _, p := range thread.Participants {
			if p != nil {
				r.Participants = append(
					r.Participants,
					participant{Username: p.Username},
				)
			}
		}

		rows = append(rows, r)
	}

	return response.SuccessAny(
		map[string]any{
			"threads":               rows,
			"total":                 threads.Total,
			"total_unread_threads":  threads.TotalUnreadThreads,
			"total_unread_mentions": threads.TotalUnreadMentions,
		},
	)
}
