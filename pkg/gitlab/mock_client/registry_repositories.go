package mock_client

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) RegistryRepositories(
	_ int64,
	_ bool,
) ([]*gitlab.RegistryRepository, error) {
	return nil, nil
}
