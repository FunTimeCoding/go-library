package client

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/filter"
	"github.com/funtimecoding/go-library/pkg/kubernetes/types/native/event"
)

func (c *Client) DeleteEventSimple(e *event.Event) {
	f := filter.New()
	f.AddClusters(e.Cluster)
	f.AddNamespaces(e.Namespace)
	f.AddNames(e.Name)
	c.DeleteEvent(f)
}
