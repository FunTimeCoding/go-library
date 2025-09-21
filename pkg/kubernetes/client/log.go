package client

import (
	"bufio"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	core "k8s.io/api/core/v1"
)

func (c *Client) Log(
	namespace string,
	pod string,
	container string,
) string {
	r := c.client.CoreV1().Pods(namespace).GetLogs(
		pod,
		&core.PodLogOptions{
			Follow:    false,
			Previous:  false,
			Container: container,
		},
	)
	l, e := r.Stream(c.context)
	errors.PanicOnError(e)
	defer errors.LogClose(l)
	s := bufio.NewScanner(l)
	var lines []string

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	errors.PanicOnError(s.Err())

	return join.NewLine(lines)
}
