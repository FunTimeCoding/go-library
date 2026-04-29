package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustReact(
	p *model.Post,
	emoji string,
) {
	errors.PanicOnError(c.React(p, emoji))
}
