package operation

import (
	"k8s.io/client-go/kubernetes"
	event "k8s.io/client-go/kubernetes/typed/events/v1"
)

func Event(
	c *kubernetes.Clientset,
	namespace string,
) event.EventInterface {
	return c.EventsV1().Events(namespace)
}
