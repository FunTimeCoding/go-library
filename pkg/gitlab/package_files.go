package gitlab

import "gitlab.com/gitlab-org/api/client-go"

func (c *Client) PackageFiles(
	project int,
	packageIdentifier int,
) []*gitlab.PackageFile {
	result, r, e := c.client.Packages.ListPackageFiles(
		project,
		packageIdentifier,
		nil,
	)
	panicOnError(r, e)

	return result
}
