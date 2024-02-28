package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) CurrentUser() *gitlab.User {
	result, r, e := c.client.Users.CurrentUser()

	if e != nil && r.StatusCode == 401 {
		// 401 is returned when using $CI_JOB_TOKEN, but this is not critical
		return nil
	}

	errors.PanicOnError(e)

	return result
}
