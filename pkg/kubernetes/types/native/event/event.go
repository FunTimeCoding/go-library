package event

import event "k8s.io/api/events/v1"

type Event struct {
	Cluster string
	Name    string
	Type    string
	Raw     *event.Event
}
