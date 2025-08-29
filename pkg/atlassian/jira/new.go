package jira

import (
	"context"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/basic_client"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/client"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/service_client"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func New(
	host string,
	user string,
	token string,
	o ...OptionFunc,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(user, "user")
	errors.FatalOnEmpty(token, "token")
	locator := fmt.Sprintf("https://%s", host)
	result := &Client{
		context: context.Background(),
		client: client.New(
			jira.BasicAuthTransport{Username: user, Password: token},
			locator,
		),
		basic:   basic_client.New(locator, user, token),
		service: service_client.New(locator, user, token),
		locator: locator,
		user:    user,
	}

	for _, p := range o {
		p(result)
	}

	if len(result.closedStatus) == 0 {
		result.closedStatus = []string{constant.Closed}
	}

	return result
}
