package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
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

	return string(system.ReadAll(l))
}
