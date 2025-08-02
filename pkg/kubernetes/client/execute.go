package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/executor"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"io"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

func (c *Client) Execute(
	namespace string,
	pod string,
	command []string,
	stdout io.Writer,
	stderr io.Writer,
) {
	r := c.client.CoreV1().RESTClient().Post().Resource(
		constant.PodsResource,
	).Namespace(namespace).Name(pod).SubResource(
		constant.ExecuteSubResource,
	)
	r.VersionedParams(
		&v1.PodExecOptions{Command: command, Stdout: true, Stderr: true},
		scheme.ParameterCodec,
	)
	errors.PanicOnError(
		executor.New(c.configuration, r).StreamWithContext(
			c.context,
			remotecommand.StreamOptions{Stdout: stdout, Stderr: stderr},
		),
	)
}
