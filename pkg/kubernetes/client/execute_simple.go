package client

import "strings"

func (c *Client) ExecuteSimple(
	namespace string,
	pod string,
	command ...string,
) (string, string) {
	stdout := &strings.Builder{}
	stderr := &strings.Builder{}
	c.Execute(stdout, stderr, namespace, pod, "", command)

	return stdout.String(), stderr.String()
}
