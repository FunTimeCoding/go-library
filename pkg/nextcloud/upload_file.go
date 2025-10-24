package nextcloud

import "github.com/funtimecoding/go-library/pkg/system"

func (c *Client) UploadFile(
	path string,
	sourcePath string,
	sourceName string,
) {
	c.authoring.WriteFile(path, system.ReadBytes(sourcePath, sourceName))
}
