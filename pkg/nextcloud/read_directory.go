package nextcloud

import "os"

func (c *Client) ReadDirectory(path string) []os.FileInfo {
	return c.authoring.ReadDirectory(path)
}
