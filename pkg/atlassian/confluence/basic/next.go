package basic

import "fmt"

func (c *Client) Next(path string) string {
	return fmt.Sprintf("https://%s%s", c.host, path)
}
