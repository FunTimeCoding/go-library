package mock_client

import (
	"encoding/base64"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) SeedFile(
	path string,
	content string,
) {
	c.files[path] = &gitlab.File{
		Content: base64.StdEncoding.EncodeToString([]byte(content)),
	}
}
