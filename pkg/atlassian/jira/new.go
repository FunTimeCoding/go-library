package jira

import (
	"context"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func New(
	host string,
	user string,
	token string,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(user, "user")
	errors.FatalOnEmpty(token, "token")
	t := jira.BasicAuthTransport{Username: user, Password: token}
	locator := fmt.Sprintf("https://%s", host)
	client, e := jira.NewClient(t.Client(), locator)
	errors.PanicOnError(e)

	return &Client{
		context: context.Background(),
		client:  client,
		locator: locator,
		user:    user,
	}
}
