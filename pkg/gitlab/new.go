package gitlab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func New(
	host string,
	token string,
	projects []int,
) *Client {
	var options []gitlab.ClientOptionFunc

	if host != "" {
		options = append(
			options,
			gitlab.WithBaseURL(fmt.Sprintf("https://%s/api/v4", host)),
		)
	}

	client, e := gitlab.NewClient(token, options...)
	errors.PanicOnError(e)
	result := &Client{client: client, projects: projects}
	result.user = result.CurrentUser()

	return result
}
