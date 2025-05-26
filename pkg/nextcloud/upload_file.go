package nextcloud

import "github.com/funtimecoding/go-library/pkg/system"

func (c *Client) UploadFile(
	path string,
	source string,
) {
	c.authoring.WriteFile(path, system.ReadBytes(source))
}
