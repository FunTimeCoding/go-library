package client

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/operation/list"
	batch "k8s.io/api/batch/v1"
)

func (c *Client) CronJobs(namespace string) []batch.CronJob {
	return list.CronJob(c.client, c.context, namespace, "")
}
