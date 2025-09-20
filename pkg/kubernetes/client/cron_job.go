package client

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/operation/get"
	"github.com/funtimecoding/go-library/pkg/kubernetes/types/native/cron_job"
)

func (c *Client) CronJob(
	namespace string,
	name string,
) *cron_job.Job {
	return cron_job.New(
		get.CronJob(c.client, c.context, namespace, name),
		c.cluster,
	)
}
