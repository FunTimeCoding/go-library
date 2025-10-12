package basic

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func New(
	host string,
	user string,
	token string,
	verbose bool,
) *Client {
	// https://developer.atlassian.com/cloud/confluence/rest/v2
	return &Client{
		host:    host,
		user:    user,
		token:   token,
		verbose: verbose,
		base:    locator.New(host).Base(constant.Base),
	}
}
