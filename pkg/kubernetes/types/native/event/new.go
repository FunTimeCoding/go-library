package event

import event "k8s.io/api/events/v1"

func New(
	v *event.Event,
	cluster string,
) *Event {
	return &Event{Cluster: cluster, Name: v.Name, Type: v.Type, Raw: v}
}
