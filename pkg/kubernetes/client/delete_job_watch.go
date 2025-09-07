package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

func (c *Client) DeleteJobWatch(
	namespace string,
	name string,
) {
	if c.Job(namespace, name) == nil {
		return
	}

	c.DeleteJob(namespace, name)
	w, e := c.client.BatchV1().Jobs(namespace).Watch(
		c.context,
		meta.ListOptions{
			FieldSelector:  key_value.Equals("metadata.name", name),
			TimeoutSeconds: ptr.To(int64(120)),
		},
	)
	errors.PanicOnError(e)
	defer w.Stop()

	for v := range w.ResultChan() {
		if v.Type == watch.Deleted {
			return
		}
	}
}
