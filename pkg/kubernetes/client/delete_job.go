package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/ptr"
	kubernetes "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) DeleteJob(
	namespace string,
	name string,
) {
	e := c.client.BatchV1().Jobs(namespace).Delete(
		c.context,
		name,
		meta.DeleteOptions{
			PropagationPolicy: ptr.To(meta.DeletePropagationForeground),
		},
	)

	if e != nil {
		if kubernetes.IsNotFound(e) {
			return
		}
	}

	errors.PanicOnError(e)
}
