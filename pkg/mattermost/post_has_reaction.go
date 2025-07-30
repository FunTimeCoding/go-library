package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) PostHasReaction(
	p *model.Post,
	reaction string,
) bool {
	if p.HasReactions {
		for _, r := range c.Reactions(p) {
			if r.EmojiName == reaction {
				return true
			}
		}
	}

	return false
}
