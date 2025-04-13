package event

import event "k8s.io/api/events/v1"

func New(
	v *event.Event,
	cluster string,
) *Event {
	return &Event{
		Cluster:       cluster,
		Namespace:     v.Namespace,
		RegardingKind: v.Regarding.Kind,
		Name:          v.Name,
		Type:          v.Type,
		Reason:        v.Reason,
		Create:        &v.CreationTimestamp.Time,
		Note:          v.Note,
		Raw:           v,
	}
}
