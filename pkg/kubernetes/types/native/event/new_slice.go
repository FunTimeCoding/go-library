package event

import event "k8s.io/api/events/v1"

func NewSlice(
	v []event.Event,
	cluster string,
) []*Event {
	var result []*Event

	for _, e := range v {
		result = append(result, New(&e, cluster))
	}

	return result
}
