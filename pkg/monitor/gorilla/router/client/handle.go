package client

import "fmt"

func (c *Client) Handle() string {
	return fmt.Sprintf("%s@%s", c.Name, c.Hostname)
}
