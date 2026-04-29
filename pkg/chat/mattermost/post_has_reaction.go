package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) PostHasReaction(
	p *model.Post,
	reaction string,
) (bool, error) {
	if p.HasReactions {
		reactions, e := c.Reactions(p)

		if e != nil {
			return false, e
		}

		for _, r := range reactions {
			if r.EmojiName == reaction {
				return true, nil
			}
		}
	}

	return false, nil
}
