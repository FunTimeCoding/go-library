package basic_client

import "github.com/funtimecoding/go-library/pkg/web/constant"

func New(
	host string,
	user string,
	token string,
	verbose bool,
) *Client {
	// https://developer.atlassian.com/cloud/confluence/rest/v2
	return &Client{
		scheme:  constant.SecureScheme,
		host:    host,
		user:    user,
		token:   token,
		verbose: verbose,
	}
}
