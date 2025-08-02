package client

import "strings"

func (c *Client) ExecuteSimple(
	namespace string,
	pod string,
	command []string,
) (string, string) {
	stdout := &strings.Builder{}
	stderr := &strings.Builder{}
	c.Execute(namespace, pod, command, stdout, stderr)

	return stdout.String(), stderr.String()
}
