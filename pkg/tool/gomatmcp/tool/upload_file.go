package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
	"os"
	"path/filepath"
)

func (t *Tool) UploadFile(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.UploadFile,
) (result *mcp.CallToolResult, e error) {
	defer func() {
		if r := recover(); r != nil {
			result = mcp.NewToolResultError(fmt.Sprintf("%v", r))
			e = nil
		}
	}()

	if a.Path == "" {
		return response.Fail("path is required")
	}

	if a.ChannelID == "" && a.ChannelName == "" {
		return response.Fail(
			"channel_id or channel_name is required",
		)
	}

	data, f := os.ReadFile(a.Path)

	if f != nil {
		return response.Fail("read file failed: %v", f)
	}

	var ch *model.Channel

	if a.ChannelName != "" {
		ch = t.client.TeamChannel(a.ChannelName)
	} else {
		ch = t.client.Channel(a.ChannelID)
	}

	upload, _, g := t.client.Nested().UploadFile(
		t.client.Context(),
		data,
		ch.Id,
		filepath.Base(a.Path),
	)

	if g != nil {
		return response.Fail("upload failed: %v", g)
	}

	if len(upload.FileInfos) == 0 {
		return response.Fail("upload returned no file info")
	}

	fileIdentifier := upload.FileInfos[0].Id
	message := a.Message

	if message == "" {
		message = filepath.Base(a.Path)
	}

	post := &model.Post{
		ChannelId: ch.Id,
		Message:   message,
		FileIds:   []string{fileIdentifier},
	}
	created, _, h := t.client.Nested().CreatePost(
		t.client.Context(),
		post,
	)

	if h != nil {
		return response.Fail("post with file failed: %v", h)
	}

	return response.SuccessAny(
		map[string]any{
			"post_id":   created.Id,
			"file_id":   fileIdentifier,
			"file_name": filepath.Base(a.Path),
			"channel":   t.channelDisplayName(ch),
			"message":   message,
			"create_at": formatMilli(created.CreateAt),
		},
	)
}
