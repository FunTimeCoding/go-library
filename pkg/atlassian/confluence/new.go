package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/kaos_client"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/virtomize_client"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func New(
	host string,
	user string,
	token string,
	labels []string,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(user, "user")
	errors.FatalOnEmpty(token, "token")

	return &Client{
		labels:    labels,
		host:      host,
		basic:     basic_client.New(host, user, token),
		kaos:      kaos_client.New(host, user, token),
		virtomize: virtomize_client.New(host, user, token),
	}
}
