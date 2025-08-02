package client

import "os"

func (c *Client) ExecuteStandard(
	namespace string,
	pod string,
	command []string,
) {
	c.Execute(namespace, pod, command, os.Stdout, os.Stderr)
}
