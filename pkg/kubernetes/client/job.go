package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	batch "k8s.io/api/batch/v1"
	kubernetes "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) Job(
	namespace string,
	name string,
) *batch.Job {
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

	return result
}
