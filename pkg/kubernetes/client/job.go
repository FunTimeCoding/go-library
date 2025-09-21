package client

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/operation/get"
	"github.com/funtimecoding/go-library/pkg/kubernetes/types/native/job"
)

func (c *Client) Job(
	namespace string,
	name string,
) *job.Job {
	result := get.Job(c.client, c.context, namespace, name)

	if result == nil {
		return nil
	}

	return job.New(result, c.cluster)
}
