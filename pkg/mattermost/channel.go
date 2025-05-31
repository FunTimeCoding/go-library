package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/mattermost/constant"
	"github.com/mattermost/mattermost-server/v6/model"
)

func (c *Client) Channel(
	t *model.Team,
	name string,
) *model.Channel {
	result, _, e := c.client.GetChannelByName(
		name,
		t.Id,
		constant.EmptyEntityTag,
	)
	errors.PanicOnError(e)

	return result
}
