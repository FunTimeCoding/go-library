package nextcloud

import (
	"github.com/funtimecoding/go-library/pkg/nextcloud/basic"
	"github.com/funtimecoding/go-library/pkg/nextcloud/helper"
	"github.com/funtimecoding/go-library/pkg/web/authoring"
)

func New(
	host string,
	user string,
	password string,
) *Client {
	return &Client{
		basic:     basic.New(host, user, password),
		authoring: authoring.New(helper.FileRoot(host, user), user, password),
	}
}
