package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/kubernetes/types/native/job"
	kubernetes "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) Job(
	namespace string,
	name string,
) *job.Job {
	result, e := c.client.BatchV1().Jobs(namespace).Get(
		c.context,
		name,
		meta.GetOptions{},
	)

	if e != nil {
		if kubernetes.IsNotFound(e) {
			return nil
		}
	}

	errors.PanicOnError(e)

	return job.New(result, c.cluster)
}
