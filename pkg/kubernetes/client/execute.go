package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/executor"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"io"
	core "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

func (c *Client) Execute(
	stdout io.Writer,
	stderr io.Writer,
	namespace string,
	pod string,
	container string,
	command ...string,
) {
	r := c.client.CoreV1().RESTClient().Post().Resource(
		constant.PodsResource,
	).Namespace(namespace).Name(pod).SubResource(
		constant.ExecuteSubResource,
	)
	r.VersionedParams(
		&core.PodExecOptions{
			Command:   command,
			Container: container,
			Stdout:    true,
			Stderr:    true,
		},
		scheme.ParameterCodec,
	)
	errors.PanicOnError(
		executor.New(c.configuration, r).StreamWithContext(
			c.context,
			remotecommand.StreamOptions{Stdout: stdout, Stderr: stderr},
		),
	)
}
