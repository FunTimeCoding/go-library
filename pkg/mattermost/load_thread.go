package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/mattermost/constant"
	"github.com/funtimecoding/go-library/pkg/mattermost/post"
	"github.com/funtimecoding/go-library/pkg/mattermost/thread"
)

func (c *Client) LoadThread(p *post.Post) *thread.Thread {
	result := thread.New(p, c.Thread(p.Raw))
	result.Resolved = c.PostHasReaction(p.Raw, constant.CheckMark)

	return result
}
