package open_webui

import "github.com/funtimecoding/go-library/pkg/generative/open_webui/basic_client"

func New(
	host string,
	token string,
) *Client {
	return &Client{basic: basic_client.New(host, token)}
}
