package client

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/operation/get"
	batch "k8s.io/api/batch/v1"
)

func (c *Client) CronJob(
	namespace string,
	name string,
) *batch.CronJob {
	return get.CronJob(c.client, c.context, namespace, name)
}
