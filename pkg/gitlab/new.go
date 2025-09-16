package gitlab

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go"
)

func New(
	host string,
	token string,
	o ...Option,
) *Client {
	result := &Client{
		context:      context.Background(),
		projectCache: make(map[int]*project.Project),
	}

	for _, p := range o {
		p(result)
	}

	var p []gitlab.ClientOptionFunc

	if host != "" {
		p = append(
			p,
			gitlab.WithBaseURL(fmt.Sprintf("https://%s/api/v4", host)),
		)
	}

	client, e := gitlab.NewClient(token, p...)
	errors.PanicOnError(e)
	result.client = client
	result.user = result.CurrentUser()

	return result
}
