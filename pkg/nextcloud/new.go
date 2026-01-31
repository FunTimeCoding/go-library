package nextcloud

import (
	"github.com/funtimecoding/go-library/pkg/nextcloud/basic"
	"github.com/funtimecoding/go-library/pkg/nextcloud/helper"
	"github.com/funtimecoding/go-library/pkg/web/author"
)

func New(
	host string,
	user string,
	password string,
) *Client {
	return &Client{
		basic:  basic.New(host, user, password),
		author: author.New(helper.FileRoot(host, user), user, password),
	}
}
