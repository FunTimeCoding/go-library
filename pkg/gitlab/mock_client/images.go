package mock_client

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) Images(
	_ int64,
	_ int64,
) ([]*gitlab.RegistryRepositoryTag, error) {
	return nil, nil
}
