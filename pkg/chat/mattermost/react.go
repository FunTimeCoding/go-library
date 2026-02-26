package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) React(
	p *model.Post,
	emoji string,
) {
	_, r, e := c.client.SaveReaction(
		c.context,
		&model.Reaction{
			UserId:    c.Me().Id,
			PostId:    p.Id,
			EmojiName: emoji,
		},
	)
	panicOnError(r, e)
}
