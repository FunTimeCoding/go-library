package gitlab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Packages(
	project int,
	panicOnForbidden bool,
) []*gitlab.Package {
	result, r, e := c.client.Packages.ListProjectPackages(
		project,
		&gitlab.ListProjectPackagesOptions{
			ListOptions: gitlab.ListOptions{PerPage: constant.PerPage1000},
		},
	)

	if !panicOnForbidden && r != nil && r.StatusCode == 403 {
		// Given correct token scope, this might be due to the GitLab server being configured wrong: https://forum.gitlab.com/t/cant-login-to-registry-due-to-denied-access-forbidden/63965/6
		fmt.Println("warning: packages 403")

		return result
	}

	errors.PanicOnError(e)

	return result
}
