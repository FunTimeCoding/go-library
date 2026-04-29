package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) PostDefault(text string) (*model.Post, error) {
	if c.channel != nil {
		return c.PostSimple(c.channel, text)
	}

	return nil, nil
}
