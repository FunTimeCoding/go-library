package mock_client

import "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/client"

func (c *Client) VersionsSince(
	_ string,
	_ int,
) []client.VersionEntry {
	return nil
}
