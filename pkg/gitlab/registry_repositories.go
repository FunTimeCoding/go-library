package gitlab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) RegistryRepositories(
	project int,
	panicOnForbidden bool,
) []*gitlab.RegistryRepository {
	result, r, e := c.client.ContainerRegistry.ListProjectRegistryRepositories(
		project,
		nil,
	)

	if !panicOnForbidden && r.StatusCode == 403 {
		// Given correct token scope, this might be due to the GitLab server being configured wrong: https://forum.gitlab.com/t/cant-login-to-registry-due-to-denied-access-forbidden/63965/6
		fmt.Println("warning: registry repositories 403")

		return result
	}

	errors.PanicOnError(e)

	return result
}
