package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustOpenDialog(v model.OpenDialogRequest) {
	errors.PanicOnError(c.OpenDialog(v))
}
