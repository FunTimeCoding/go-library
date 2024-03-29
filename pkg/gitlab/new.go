package gitlab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func New(
	host string,
	token string,
) *Client {
	client, e := gitlab.NewClient(
		token,
		gitlab.WithBaseURL(fmt.Sprintf("https://%s/api/v4", host)),
	)
	errors.PanicOnError(e)
	result := &Client{client: client}
	result.user = result.CurrentUser()

	return result
}
